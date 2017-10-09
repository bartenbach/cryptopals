package challenge2

import (
	"bytes"
	"com/blakebartenbach/cryptopals/challenge1"
	"log"
	"testing"
)

func TestProblem2(t *testing.T) {
	// decode all input values from hex
	val1 := challenge1.DecodeHex("1c0111001f010100061a024b53535009181c")
	val2 := challenge1.DecodeHex("686974207468652062756c6c277320657965")
	val3 := challenge1.DecodeHex("746865206b696420646f6e277420706c6179")

	// print these out because of cryptopals easter eggs
	log.Println(string(val1))
	log.Println(string(val2))
	log.Println(string(val3))

	// XOR them
	res := XOR(val1, val2)

	// check to see if the xored values produced the expected result (val3) and log results
	if !bytes.Equal(res, val3) {
		log.Fatal("Incorrect result!  Problem2 failed")
	} else {
		log.Println("Problem 2 passed!")
	}
}
