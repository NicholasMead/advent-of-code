package mockfs

import "testing"

func TestDirectory(t *testing.T) {
	t.Run("addFile", func(t *testing.T) {
		dir := CreateDirectory("dir")

		dir.AddFile("a", 5)

		if size, err := dir.GetFileSize("a"); err != nil {
			t.Error(err)
		} else if size != 5 {
			t.Error("expected file of size of 5 got", size)
		}
	})

	t.Run("addDirectory", func(t *testing.T) {
		dir := CreateDirectory("dir")
		sub := CreateDirectory("sub")

		dir.AddDirectory(sub)

		if ret, err := dir.GetSubDirectory("sub"); err != nil {
			t.Error(err)
		} else if name := ret.GetName(); name != "sub" {
			t.Error("expected directory with name sub got", name)
		} else if ret != sub {
			t.Error("Different dir returned")
		}
	})

	t.Run("getSize", func(t *testing.T) {
		dir := CreateDirectory("dir")
		dir.AddFile("a", 1)

		sub1 := CreateDirectory("sub1")
		sub1.AddFile("a", 1)

		sub2 := CreateDirectory("sub2")
		sub2.AddFile("a", 10)
		sub2.AddFile("b", 20)

		sub21 := CreateDirectory("sub21")
		sub21.AddFile("a", 100)
		sub21.AddFile("b", 200)
		sub21.AddFile("c", 300)

		dir.AddDirectory(sub1)
		dir.AddDirectory(sub2)
		sub2.AddDirectory(sub21)

		if size := dir.GetSize(); size != 632 {
			t.Error("expected directory size of 632 got", size)
		}
	})
}
