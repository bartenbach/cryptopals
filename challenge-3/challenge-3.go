package challenge_3

import (
	"../challenge-1"
	"../challenge-2"
	"bytes"
	"fmt"
)

func Challenge3() {
	// get encoded hex value and decode
	var hex = challenge_1.GetString()
	var decoded = challenge_1.DecodeHex(hex)

	// would you believe I guessed this immediately?
	s0 := bytes.Repeat([]byte{'X'}, 34)
	var xored = challenge_2.XORvalues(decoded, s0)
	fmt.Println(string(xored))
}
