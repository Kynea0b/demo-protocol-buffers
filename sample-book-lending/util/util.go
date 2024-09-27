package util

import (
	"encoding/binary"
)

func EncodeUint(x uint64) ([]byte, int) {
	buf := make([]byte, binary.MaxVarintLen64)

	n := binary.PutUvarint(buf, x)

	return buf[:n], n
}

func DecodeUint(buf []byte) uint64 {
	val, _ := binary.Uvarint(buf[:])
	return uint64(val)
}
