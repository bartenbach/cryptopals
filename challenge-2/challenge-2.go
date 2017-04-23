package challenge_2

import (
	"com/blakebartenbach/cryptopals/challenge-1"
	"encoding/hex"
	"fmt"
	"log"
)

func Challenge2() {
	// get first hex value and decode
	var input string = challenge_1.GetString()
	var decoded = challenge_1.DecodeHex(input)

	// get second hex value and decode
	var input2 string = challenge_1.GetString()
	var decoded2 = challenge_1.DecodeHex(input2)

	// XOR against each other and encode to hex
	var xored, err = XORvalues(decoded, decoded2)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(hex.EncodeToString(xored))
	}
}

func XORvalues(value1, value2 []byte) ([]byte, error) {
	if len(value1) != len(value2) {
		return nil, fmt.Errorf("Lengths of byte slices are not equal! %d != %d", len(value1), len(value2))
	}

	buffer := make([]byte, len(value1))

	for i := range value1 {
		buffer[i] = value1[i] ^ value2[i]
	}

	return buffer, nil
}
