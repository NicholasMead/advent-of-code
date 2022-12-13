package packet

import (
	"fmt"
	"strings"
)

type ArrayPacket interface {
	Packet
	Array() []Packet
	Value(i int) Packet
}

func CreatePacketArray(v ...Packet) ArrayPacket {
	return &arrayPacket{v}
}

func CreateValueArray(v ...int) ArrayPacket {
	values := make([]Packet, len(v))
	for i := range v {
		values[i] = CreateValuePacket(v[i])
	}
	return &arrayPacket{values}
}

type arrayPacket struct {
	value []Packet
}

// Copy implements ArrayPacket
func (a *arrayPacket) DeepCopy() Packet {
	packets := make([]Packet, len(a.value))
	for i, p := range a.value {
		packets[i] = p.DeepCopy()
	}
	return &arrayPacket{packets}
}

// GetType implements ValuePacket
func (*arrayPacket) GetType() PacketType {
	return array
}

// Value implements ValuePacket
func (a *arrayPacket) Array() []Packet {
	v := make([]Packet, len(a.value))
	copy(v, a.value)
	return v
}

func (a *arrayPacket) Value(i int) Packet {
	return a.value[i]
}

func (a *arrayPacket) String() string {
	elems := []string{}
	for _, p := range a.value {
		elems = append(elems, fmt.Sprintf("%v", p))
	}
	return fmt.Sprintf("[%v]", strings.Join(elems, ","))
}
