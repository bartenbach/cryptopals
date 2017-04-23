package challenge_6

import (
    "com/blakebartenbach/cryptopals/challenge-4"
    "fmt"
    "encoding/base64"
    "com/blakebartenbach/cryptopals/challenge-1"
)

func Challenge6() {
    var elements []challenge_4.EncodedElement = challenge_4.GetEncodedElements("challenge-6/6.txt")
    for i := range elements {
        data, err := base64.StdEncoding.DecodeString(elements[i].Value)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(data)
        }
    }

    var input1 string = challenge_1.GetString()
    var input2 string = challenge_1.GetString()

}

func HammingDistance(x,y string) {
    
}

