package challenge_5

import "testing"

func TestGetRepeatingXOR(t *testing.T) {
    y := "ICE"
    x := []byte{ 'I', 'C', 'E', 'I'}
    xor := GetRepeatingXOR([]byte(y), 4)
    if string(x) != string(xor) {
       t.Error("Repeating XOR Failed")
    }

}
