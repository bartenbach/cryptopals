package challenge_6

import (
	"com/blakebartenbach/cryptopals/challenge-4"
	"encoding/base64"
	"fmt"
)

var table = [256]uint8{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

func Challenge6() {
	// Read in elements from encrypted file - store them in a slice of EncodedElements (really just strings)
	var elements []challenge_4.EncodedElement = challenge_4.GetEncodedElements("challenge-6/6.txt")
	for i := range elements {
		data, err := base64.StdEncoding.DecodeString(elements[i].Value)
		if err != nil {
			fmt.Println(err)
		} else {
			PrintPrettySlice(data)
		}
	}

	//Guess our keysize.  Will have to find the keysize with the shortest hamming distance.
	// should I do this manually or programmatically?
	keysize := 4

	// Using our keysize, calculate hamming distance between two keys
	for j := 0; j < len(elements); j += keysize {
		// need to get the first <keysize> worth of bytes

		// wow...this is confusing
		bytes1 := elements[j : j+keysize]
		bytes2 := elements[j+keysize : j+(keysize*2)]
		hammingd, err := HammingDistance(bytes1, bytes2)
		if err != nil {
			println("ERROR: ", err)
		} else {
			// normalize hamming distance by dividing by the keysize
			fmt.Printf("Hamming distance: %d\n", hammingd/keysize)
		}

	}

}

func HammingDistance(x, y []byte) (int, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("String lengths not equal! %d != %d", len(x), len(y))
	}

	b1 := []byte(x)
	b2 := []byte(y)

	var hammingd int = 0

	for i := range b1 {
		hammingd += int(table[(b1[i] ^ b2[i])])
	}
	return hammingd, nil
}

// Uses printf to print evenly spaced columns for slices
func PrintPrettySlice(b []byte) {
	fmt.Printf("[")
	for x := range b {
		fmt.Printf("%5d", b[x])
	}
	fmt.Printf("]\n")
}
