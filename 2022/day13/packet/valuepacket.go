package packet

import "fmt"

type ValuePacket interface {
	Packet
	Value() int
	ToArray() ArrayPacket
}

func CreateValuePacket(v int) ValuePacket {
	return &valuePacket{v}
}

type valuePacket struct {
	value int
}

// Copy implements ValuePacket
func (v *valuePacket) DeepCopy() Packet {
	value := v.value
	return &valuePacket{value}
}

// ToArray implements ValuePacket
func (v *valuePacket) ToArray() ArrayPacket {
	return &arrayPacket{[]Packet{v}}
}

// GetType implements ValuePacket
func (*valuePacket) GetType() PacketType {
	return value
}

// value implements ValuePacket
func (v *valuePacket) Value() int {
	return v.value
}

func (v *valuePacket) String() string {
	return fmt.Sprint(v.value)
}
