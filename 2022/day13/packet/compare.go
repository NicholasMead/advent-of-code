package packet

import "fmt"

func Ordered(left Packet, right Packet) (inOrder bool) {
	return Compare(left, right) >= 0
}

func Compare(left Packet, right Packet) int {
	return valueCompare(left, right)
}

func valueCompare(left Packet, right Packet) (inOrder int) {
	lType, rType := left.GetType(), right.GetType()

	//Both value type
	switch {
	case lType == value && rType == value:
		lVal, rVal := left.(ValuePacket), right.(ValuePacket)
		return rVal.Value() - lVal.Value()

	case lType == array && rType == array:
		lVal, rVal := left.(ArrayPacket), right.(ArrayPacket)
		return arrayCompare(lVal, rVal)

	case lType == value && rType == array:
		lVal, rVal := CreatePacketArray(left), right.(ArrayPacket)
		return arrayCompare(lVal, rVal)

	case lType == array && rType == value:
		lVal, rVal := left.(ArrayPacket), CreatePacketArray(right)
		return arrayCompare(lVal, rVal)

	default:
		err := fmt.Sprintf("Unknown types %v & %v", lType, rType)
		panic(err)
	}
}

func arrayCompare(left ArrayPacket, right ArrayPacket) (inOrder int) {
	index := 0

	for {
		if len(left.Array()) <= index || len(right.Array()) <= index {
			return len(right.Array()) - len(left.Array())
		}

		valResult := valueCompare(left.Value(index), right.Value(index))
		if valResult == 0 {
			index++
		} else {
			return valResult
		}
	}
}
