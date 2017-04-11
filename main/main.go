package main

import "fmt"
import "../challenge-1"

func main() {
    challenge2()
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

}