package challenge_2

import (
    "github.com/hashicorp/vault/helper/xor"
    "log"
)

func XORvalues(value1, value2 []byte) []byte  {
    var xored, err = xor.XORBytes(value1,value2)
    if err != nil {
        log.Fatal(err)
        return nil
    } else {
        return xored
    }
}