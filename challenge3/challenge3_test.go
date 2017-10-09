package challenge3

import "testing"
import "io/ioutil"
import "log"

import "com/blakebartenbach/cryptopals/challenge1"

func TestProblem3(t *testing.T) {
	// build corpus
	text, err := ioutil.ReadFile("theodyssey.txt")
	if err != nil {
		t.Fatal("Failed to read input text:", err)
	}
	corpus := BuildCorpus(string(text))

	// decode input string and log to console
	input := challenge1.DecodeHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	result := FindSingleXORKey(input, corpus)
	log.Println(string(result))
}
