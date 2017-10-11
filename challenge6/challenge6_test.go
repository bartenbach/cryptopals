package challenge6

import "testing"

import "strings"
import "log"
import "com/blakebartenbach/cryptopals/challenge3"

func TestCalculateHammingDistance(t *testing.T) {
	if CalculateHammingDistance([]byte("this is a test"), []byte("wokka wokka!!!")) != 37 {
		t.Fatal("Hamming Distance function failure")
	}
}

func TestProblem6(t *testing.T) {
	// read in the challnge file
	file := ReadFile("6.txt")

	// remove all newlines in the file
	trimmedFile := strings.Replace(string(file), "\n", "", -1)

	// decode the file
	decodedFile := DecodeBase64(trimmedFile)

	// get the size of our repeating XOR key
	keySize := GetRepeatingXORSize(decodedFile, 2, 40)

	// prepare the corpus
	corpus := challenge3.GetCorpusFromFile("../challenge3/theodyssey.txt")

	log.Printf("%.5f", corpus['A'])

	log.Println("Likely key: ", keySize)
}
