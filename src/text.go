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

func obtainStartOfText() string {
	var result string
	for i := 0; i < gramNum-1; i++ {
		result += string(rune(2))
	}
	return result
}

func obtainEndOfText() string {
	var result string
	for i := 0; i < gramNum-1; i++ {
		result += string(rune(3))
	}
	return result
}
