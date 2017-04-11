package main

import "fmt"
import "../challenge-1"

func main() {

    var input string = challenge_1.GetString()

    var decoded = challenge_1.DecodeHex(input)

    

    fmt.Println(decoded)
}
