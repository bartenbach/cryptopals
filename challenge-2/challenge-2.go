package main

import "fmt"

func main() {

    var input string = getString()

    var decoded = decodeHex(input)

    fmt.Println(decoded)
}
