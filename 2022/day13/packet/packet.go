package packet

import "fmt"

type Packet interface {
	GetType() PacketType
	DeepCopy() Packet
}

type PacketType string

const (
	value PacketType = "value"
	array PacketType = "array"
)

func Equals(a, b Packet) bool {
	return fmt.Sprint(a) == fmt.Sprint(b)
}
