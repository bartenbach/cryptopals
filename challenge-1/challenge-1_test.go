package challenge_1

import "testing"

func TestDecodeHex(t *testing.T) {
    val := DecodeHex("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
    if string(val) != "I'm killing your brain like a poisonous mushroom" {
        t.Error("DecodeHex failed")
    }
}

func TestEncodeBase64(t *testing.T) {
    val := EncodeBase64([]byte("I'm killing your brain like a poisonous mushroom"))
    if val != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
        t.Error("EncodeBase64 failed")
    }
}
