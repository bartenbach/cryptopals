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

    //Guess our keysize.  I have no fucking clue?
    keysize := 2

    // Using our keysize, calculate hamming distance between two keys
    for j := range elements {
        // need to get the first <keysize> worth of bytes

    }

}

func HammingDistance(x, y string) (int, error) {
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
