package main

import (
    "../challenge-1"
    "../challenge-2"
    "fmt"
    "encoding/hex"
    "bytes"
)

// proof of completing challenges can be easily uncommented and proven here
func main() {
    //challenge1()
    //challenge2()
    //challenge3()
}

func challenge1() {
    // get input string
    var value string = challenge_1.GetString()

    // decode hex
    var decoded []byte = challenge_1.DecodeHex(value)

    // encode base64
    var encoded = challenge_1.EncodeBase64(decoded)

    // print result
    fmt.Println(string(encoded))
}

func challenge2() {
    // get first hex value and decode
    var input string = challenge_1.GetString()
    var decoded = challenge_1.DecodeHex(input)

    // get second hex value and decode
    var input2 string = challenge_1.GetString()
    var decoded2 = challenge_1.DecodeHex(input2)

    // XOR against each other and encode to hex
    var xored = challenge_2.XORvalues(decoded, decoded2)
    fmt.Println(hex.EncodeToString(xored))
}

func challenge3() {
    // get encoded hex value and decode
    var hex = challenge_1.GetString()
    var decoded = challenge_1.DecodeHex(hex)

    // would you believe I guessed this immediately?
    s0 := bytes.Repeat([]byte{ 'X' }, 34)
    var xored = challenge_2.XORvalues(decoded, s0)
    fmt.Println(string(xored))
}