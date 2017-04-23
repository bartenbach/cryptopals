package challenge_6

import (
	"com/blakebartenbach/cryptopals/challenge-1"
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
	var elements []challenge_4.EncodedElement = challenge_4.GetEncodedElements("challenge-6/6.txt")
	for i := range elements {
		data, err := base64.StdEncoding.DecodeString(elements[i].Value)
		if err != nil {
			fmt.Println(err)
		} else {
			PrintPrettySlice(data)
		}
	}

	var input1 string = challenge_1.GetString()
	var input2 string = challenge_1.GetString()
	hammingd, err := HammingDistance(input1, input2)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println("Hamming distance: ", hammingd)
	}

}

func HammingDistance(x, y string) (int, error) {
	// I think these have to be the same size
	// What can I return here?
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
