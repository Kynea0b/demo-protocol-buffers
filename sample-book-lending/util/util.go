package util

import (
	// "fmt"
	"encoding/binary"
	"strings"
	"time"
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

var layout = "2006-01-02 15:04:05"

func TimeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

func StringToTime(str string) time.Time {
	strs := strings.Split(str, ".")
	t, _ := time.Parse("2006-01-02 15:04:05", strs[0])
	return t
}
