package challenge_3

import (
    "com/blakebartenbach/cryptopals/challenge-1"
    "bytes"
    "com/blakebartenbach/cryptopals/challenge-2"
    "fmt"
)

func Challenge3() {
    // get encoded hex value and decode
    var hex = challenge_1.GetString()
    var decoded = challenge_1.DecodeHex(hex)

    // would you believe I guessed this immediately?
    s0 := bytes.Repeat([]byte{ 'X' }, 34)
    var xored = challenge_2.XORvalues(decoded, s0)
    fmt.Println(string(xored))
}
