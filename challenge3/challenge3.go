package challenge3

import (
	"math"
	"unicode/utf8"
)

// BuildCorpus uses a popular English text (Homer's The Odyssey),
// to build a map of character frequency for English.  For any given
// rune (character), there will be a corresponding frequency in the
// form of a float64.  I would expect values like 'a' and 'e' to have
// relatively high values compared to values like 'x' or 'z'
func BuildCorpus(input string) map[rune]float64 {
	corpus := make(map[rune]float64)
	for _, char := range input {
		corpus[char]++
	}
	total := utf8.RuneCountInString(input)
	for char := range corpus {
		corpus[char] = corpus[char] / float64(total)
	}
	return corpus
}

// ScoreText scores a string of text based on the input corpus.  Higher scores
// will come from English text, much lower scores from gibberish containing
// non-printable characters.
func ScoreText(input string, corpus map[rune]float64) float64 {
	var score float64
	for _, char := range input {
		score += corpus[char]
	}
	return score / float64(utf8.RuneCountInString(input))
}

// SingleCharXOR performs an XOR on the input array of bytes using the
// provided input key.
func SingleCharXOR(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i, char := range input {
		result[i] = char ^ key
	}
	return result
}

// FindSingleXORKey attempts to find the key to a single character XOR by
// testing all possible byte values and scoring them against the
// corpus.  The highest scoring value will be returned as an array of
// bytes
func FindSingleXORKey(input []byte, corpus map[rune]float64) []byte {
	var bestScore float64
	var result []byte
	for key := byte(0); key < math.MaxUint8; key++ {
		xored := SingleCharXOR(input, key)
		score := ScoreText(string(xored), corpus)
		if score > bestScore {
			result = xored
			bestScore = score
		}
	}
	return result
}
