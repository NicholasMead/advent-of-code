package packet

import "testing"

func TestCompare(t *testing.T) {
	t.Run("Value-Value", func(t *testing.T) {
		for _, test := range []struct {
			left, right Packet
			expect      bool
		}{
			{CreateValuePacket(1), CreateValuePacket(1), true},
			{CreateValuePacket(1), CreateValuePacket(2), true},
			{CreateValuePacket(2), CreateValuePacket(1), false},
		} {
			result := Ordered(test.left, test.right)
			if result != test.expect {
				t.Errorf("Got %v expected %v for l:%v, r:%v", result, test.expect, test.left, test.right)
			}
		}
	})

	t.Run("Array-Array", func(t *testing.T) {
		for _, test := range []struct {
			left, right Packet
			expect      bool
		}{
			{CreateValueArray(1, 1, 3), CreateValueArray(1, 1, 5), true},
			{CreateValueArray(2, 3, 4), CreateValueArray(4), true},
			{CreateValueArray(9), CreateValueArray(8, 7, 6), false},
			{CreateValueArray(7, 7, 7, 7), CreateValueArray(7, 7, 7), false},
			{CreateValueArray(7, 7, 7), CreateValueArray(7, 7, 7, 7), true},
		} {
			result := Ordered(test.left, test.right)
			if result != test.expect {
				t.Errorf("Got %v expected %v for l:%v, r:%v", result, test.expect, test.left, test.right)
			}
		}
	})

	t.Run("Mixed", func(t *testing.T) {
		for _, test := range []struct {
			left, right Packet
			expect      bool
		}{
			{CreateValueArray(1), CreateValuePacket(1), true},
			{CreateValuePacket(4), CreateValueArray(2), false},
		} {
			result := Ordered(test.left, test.right)
			if result != test.expect {
				t.Errorf("Got %v expected %v for l:%v, r:%v", result, test.expect, test.left, test.right)
			}
		}
	})

}
