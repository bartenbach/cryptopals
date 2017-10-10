package challenge5

import (
	"bytes"
	"com/blakebartenbach/cryptopals/challenge1"
	"testing"
)

func TestGetRepeatingXORArray(t *testing.T) {
	// define input string
	input := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)

	// do repeating xor of the input with the required array
	result := RepeatingXOR(input, []byte{'I', 'C', 'E'})

	// define expected output
	output := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	// fail if result is not expected output
	if !bytes.Equal(result, challenge1.DecodeHex(output)) {
		t.Fatal("Repeating XOR failed - byte arrays not equal!")
	}

}
