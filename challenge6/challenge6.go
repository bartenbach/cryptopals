package challenge6

import (
	"com/blakebartenbach/cryptopals/challenge3"
	"encoding/base64"
	"io/ioutil"
	"math"
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

// GetRepeatingXORSize attempts to guess the key to something that has
// been encrypted via repeating XOR.  It guesses key sizes from min to max,
// returning the key size that had the lowest hamming distance.
func GetRepeatingXORSize(input []byte, min int, max int) int {
	var smallestKey float32
	smallestKey = math.MaxFloat32
	var likelyKey int
	multiplier := 8
	for i := min; i < max; i++ {
		chunk1 := input[0 : i*multiplier]
		chunk2 := input[i*multiplier : i*2*multiplier]
		hammingd := float32(CalculateHammingDistance(chunk1, chunk2)) / float32(i)
		if hammingd < smallestKey { // this is really confusing is this right?
			smallestKey = hammingd
			likelyKey = i
		}
	}
	return likelyKey
}

// GetRepeatingXORKey takes the input with the likely keysize, and attempts
// to derive the key (which it then returns)
func GetRepeatingXORKey(input []byte, keySize int, corpus map[rune]float64) []byte {
	column := make([]byte, len(input)/keySize)
	key := make([]byte, keySize)
	for col := 0; col < keySize; col++ {
		for i := range column {
			column[i] = input[i*keySize+col]
		}
		_, _, k := challenge3.FindSingleXORKey(column, corpus)
		key[col] = k
	}
	return key
}

// DecodeBase64 decodes an input string from base64
func DecodeBase64(input string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		panic("Failed to decode base64")
	}
	return decoded
}

// ReadFile reads a file of the provided name and returns the bytes
func ReadFile(fileName string) []byte {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("Failed to read file!")
	}
	return file
}
