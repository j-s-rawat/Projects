package main

import (
	"encoding/binary"
)

func byteArrayToInt(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	var value int
	if data[0]&0x80 != 0 { // Check the sign bit
		a := []byte{0, 0, 0, 0xfe} //append([]byte{0, 0, 0, 0}, data...)
		value = int(binary.BigEndian.Uint32(a))
	} else {
		value = int(binary.BigEndian.Uint32(data))
	}

	return value
}

/* func main() {
	// Example usage
	byteArray := []byte{0xFF, 0xFF, 0xFF, 0xFE}
	result := byteArrayToInt(byteArray)
	fmt.Println(result) // Output: -2
} */
