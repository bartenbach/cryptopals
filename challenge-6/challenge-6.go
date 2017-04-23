package challenge_6

import (
	"com/blakebartenbach/cryptopals/challenge-1"
	"com/blakebartenbach/cryptopals/challenge-4"
	"encoding/base64"
	"fmt"
	"github.com/steakknife/hamming"
)

func Challenge6() {
	var elements []challenge_4.EncodedElement = challenge_4.GetEncodedElements("challenge-6/6.txt")
	for i := range elements {
		data, err := base64.StdEncoding.DecodeString(elements[i].Value)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
	}

	var input1 string = challenge_1.GetString()
	var input2 string = challenge_1.GetString()
	hammingd := HammingDistance(input1, input2)

	fmt.Println("Distance: ", hammingd)

}

func HammingDistance(x, y string) int {
	// I think these have to be the same size
	// What can I return here?
	if len(x) != len(y) {
		return 0
	}

	b1 := []byte(x)
	b2 := []byte(y)

	var hammingd int = 0

	for i := range b1 {
		hammingd += hamming.Byte(b1[i], b2[i])
	}
	return hammingd
}
