package challenge3

import (
	"io/ioutil"
	"math"
	"unicode/utf8"
)

func buildCorpus(input string) map[rune]float64 {
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

// GetCorpusFromFile returns a corpus map from a given file path.
// The map will map characters (runes) to their frequency in the text (float64).
// For any given character (rune), there will be a corresponding
// frequency contained in the map, that corresponds to the frequency
// of that character within the provided text (usually a book). In
// English text, characters like 'a' and 'i' will have relatively high
// frequencies whereas values like 'x' and 'z' will be relatively low.
func GetCorpusFromFile(fpath string) map[rune]float64 {
	corpusFile, err := ioutil.ReadFile(fpath)
	if err != nil {
		panic("Failed to read corpus file!")
	}
	return buildCorpus(string(corpusFile))
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
// bytes, along with the best score, and the key
func FindSingleXORKey(input []byte, corpus map[rune]float64) ([]byte, float64, byte) {
	var bestScore float64
	var result []byte
	var bestKey byte
	for key := byte(0); key < math.MaxUint8; key++ {
		xored := SingleCharXOR(input, key)
		score := ScoreText(string(xored), corpus)
		if score > bestScore {
			result = xored
			bestScore = score
			bestKey = key
		}
	}
	return result, bestScore, bestKey
}
