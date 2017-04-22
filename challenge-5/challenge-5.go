package challenge_5

import (
	"com/blakebartenbach/cryptopals/challenge-1"
	"com/blakebartenbach/cryptopals/challenge-2"
	"encoding/hex"
	"fmt"
)

func Challenge5() {
	// get the string we will be XORing
	var input string = challenge_1.GetString()
	var binput []byte = []byte(input)
	var binputLength int = len(binput)

	// get the XOR sequence
	var xor string = challenge_1.GetString()
	var bxor []byte = []byte(xor)

	// create repeating-key XOR
	var repeatingXORkey []byte = GetRepeatingXOR(bxor, binputLength)

	// XOR the input string and the repeating-key
	var XORed []byte = challenge_2.XORvalues(binput, repeatingXORkey)
	var encoded string = hex.EncodeToString(XORed)
	fmt.Println(encoded)
}

// Accepts byte array of XOR data & needs length of string that it will be XORed against
func GetRepeatingXOR(xor []byte, length int) []byte {
	s1 := make([]byte, length)
	var j int = 0
	for i := 0; i < length; i++ {
		s1[i] = xor[j]
		if j >= (len(xor) - 1) {
			j = 0
		} else {
			j++
		}
	}
	return s1
}
