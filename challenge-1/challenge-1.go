package main

import (
    "fmt"
    "log"
    "encoding/hex"
    "bufio"
    "os"
    "encoding/base64"
)

func main() {
    // get input string
    var value string = getString()

    // decode hex
    var decoded []byte = decodeHex(value)

    // encode base64
    var encoded = encodeBase64(decoded)

    // print result
    fmt.Println(string(encoded))
}

func getString() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the string to decode: ")
    value, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("could not read string!")
        return ""
    }
    // this line is extremely important as it trims the newline byte.  will not decode without this line
    return value[:len(value) - 1]
}

func decodeHex(value string) []byte {
    decoded, err := hex.DecodeString(value)
    if err != nil {
        log.Fatal(err)
        return nil
    } else {
        return decoded
    }
}

func encodeBase64(value []byte) string {
    return base64.URLEncoding.EncodeToString(value)
}
