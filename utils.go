package core

import (
	"bytes"
	"encoding/binary"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	_ = binary.Write(buff, binary.BigEndian, num)
	return buff.Bytes()
}