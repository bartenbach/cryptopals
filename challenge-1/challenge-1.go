package challenge_1

import (
    "fmt"
    "log"
    "encoding/hex"
    "bufio"
    "os"
    "encoding/base64"
)

func GetString() string {
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

func DecodeHex(value string) []byte {
    decoded, err := hex.DecodeString(value)
    if err != nil {
        log.Fatal(err)
        return nil
    } else {
        return decoded
    }
}

func EncodeBase64(value []byte) string {
    return base64.URLEncoding.EncodeToString(value)
}
