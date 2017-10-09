package challenge1

import "testing"
import "log"

func TestProblem1(t *testing.T) {
	// decode the hex input value
	val := DecodeHex("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	// log it to console because easter egg
	log.Println(string(val))

	// encode it as base64
	val2 := EncodeBase64(val)

	// verify result is the expected string to solve the problem
	if string(val2) != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		log.Fatal("Problem 1 failed: did not get expected result!")
	} else {
		log.Println("Problem 1 passed!")
	}
}

func TestDecodeHex(t *testing.T) {
	val := DecodeHex("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if string(val) != "I'm killing your brain like a poisonous mushroom" {
		t.Fatal("DecodeHex failed")
	}
}

func TestEncodeBase64(t *testing.T) {
	val := EncodeBase64([]byte("I'm killing your brain like a poisonous mushroom"))
	if val != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Fatal("EncodeBase64 failed")
	}
}
