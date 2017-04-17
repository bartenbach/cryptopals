package challenge_6

import (
    "com/blakebartenbach/cryptopals/challenge-4"
    "fmt"
)

func Challenge6() {
    var elements []challenge_4.EncodedElement = challenge_4.GetEncodedElements("challenge-6/6.txt")
    for _, Element := range elements {
        fmt.Println(Element)
    }
}