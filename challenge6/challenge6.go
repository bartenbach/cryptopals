package challenge6

import (
	"math/bits"
)

// CalculateHammingDistance returns the hamming distance between two
// strings (or arrays of bytes)
func CalculateHammingDistance(val1, val2 []byte) int {
	if len(val1) != len(val2) {
		panic("Different lengths!")
	}
	var result int
	for i := range val1 {
		result += bits.OnesCount8(val1[i] ^ val2[i])
	}
	return result
}