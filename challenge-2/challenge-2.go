package challenge_2

import (
    "github.com/hashicorp/vault/helper/xor"
    "log"
    "com/blakebartenbach/cryptopals/challenge-1"
    "fmt"
    "encoding/hex"
)

func Challenge2() {
     // get first hex value and decode
    var input string = challenge_1.GetString()
    var decoded = challenge_1.DecodeHex(input)

    // get second hex value and decode
    var input2 string = challenge_1.GetString()
    var decoded2 = challenge_1.DecodeHex(input2)

    // XOR against each other and encode to hex
    var xored = XORvalues(decoded, decoded2)
    fmt.Println(hex.EncodeToString(xored))
}

func XORvalues(value1, value2 []byte) []byte  {
    var xored, err = xor.XORBytes(value1,value2)
    if err != nil {
        log.Fatal(err)
        return nil
    } else {
        return xored
    }
}