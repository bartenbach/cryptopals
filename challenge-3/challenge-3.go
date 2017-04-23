package challenge_3

import (
	"bytes"
	"com/blakebartenbach/cryptopals/challenge-1"
	"com/blakebartenbach/cryptopals/challenge-2"
	"fmt"
	"log"
)

func Challenge3() {
	// get encoded hex value and decode
	var hex = challenge_1.GetString()
	var decoded = challenge_1.DecodeHex(hex)

	// would you believe I guessed this immediately?
	s0 := bytes.Repeat([]byte{'X'}, 34)
	var xored, err = challenge_2.XORvalues(decoded, s0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(xored))
	}
}
