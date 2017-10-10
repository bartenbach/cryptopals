package challenge4

import (
	"com/blakebartenbach/cryptopals/challenge1"
	"com/blakebartenbach/cryptopals/challenge3"
	"io/ioutil"
	"strings"
	"testing"
)

func TestProblem4(t *testing.T) {
	// setup corpus
	corpus := challenge3.GetCorpusFromFile("../challenge3/theodyssey.txt")

	// read in challenge file
	file, err := ioutil.ReadFile("4.txt")
	if err != nil {
		t.Fatal("Failed to read challenge file:", err)
	}
	var bestScore float64
	var result []byte
	for _, line := range strings.Split(string(file), "\n") {
		decoded := challenge1.DecodeHex(line)
		value, score := challenge3.FindSingleXORKey(decoded, corpus)
		if score > bestScore {
			bestScore = score
			result = value
		}
	}
	t.Log(string(result))
}
