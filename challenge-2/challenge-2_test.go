package challenge_2

import (
	"com/blakebartenbach/cryptopals/challenge-2"
	"encoding/hex"
	"testing"
)

func TestXORvalues(t *testing.T) { //c
	val1, err1 := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	val2, err2 := hex.DecodeString("686974207468652062756c6c277320657965")
	result, err3 := challenge_2.XORvalues([]byte(val1), []byte(val2))

	if err1 != nil || err2 != nil || err3 != nil {
		t.Fail()
	}

	if hex.EncodeToString(result) != "746865206b696420646f6e277420706c6179" {
		t.Errorf("XOR Values Failed!\nExpected '746865206b696420646f6e277420706c6179' got %s", hex.EncodeToString(result))
		t.Fail()
	}
}
