package main

import "encoding/binary"

func intToByteArray(num int) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(num))
	return b
}

func byteArrayToInt(b []byte) int {
	return int(binary.LittleEndian.Uint32(b))
}
