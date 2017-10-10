package challenge3

import (
	"com/blakebartenbach/cryptopals/challenge1"
	"testing"
)

func TestProblem3(t *testing.T) {
	// get corpus
	corpus := GetCorpusFromFile("theodyssey.txt")

	// decode input string from challenge
	input := challenge1.DecodeHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	// find the key and log to console
	result, _ := FindSingleXORKey(input, corpus)
	t.Log(string(result))
}
