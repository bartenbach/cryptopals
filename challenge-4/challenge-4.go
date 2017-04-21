package challenge_4

import (
	"com/blakebartenbach/cryptopals/challenge-1"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"com/blakebartenbach/cryptopals/challenge-2"
)

type EncodedElement struct {
	Value string
}

func Challenge4() {
	// create slice of elements from file
	var elements []EncodedElement = GetEncodedElements("challenge-4/4.txt")

	// create alphabet array of bytes to XOR against
alphabet := [68]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1',
		'2', '3', '4', '5', '6', '7', '8', '9', '~', '!', '@', '#', '$', '%', '^', '&',
		'*', '(', ')', '-', '_', '=', '+', '{', '}', '[', ']', '\\', '|', '/', '.', '<',
		'>', ',', '`', '?', '\'', '\'', ':', ';'}
	// loop through the entire alphabet testing
	for i := range alphabet {
		// get our current character
		var letter = alphabet[i]
		fmt.Println("Current letter " + string(letter))

		for j := range elements {
			var decoded []byte = challenge_1.DecodeHex(elements[j].Value)
			s0 := bytes.Repeat([]byte{letter}, 30)
			var xored = challenge_2.XORvalues(decoded, s0)
			if strings.Contains(string(xored), "the") {
				fmt.Println(string(xored))
			}
		}
	}
}

func GetEncodedElements(fpath string) []EncodedElement {
	path, err := filepath.Abs(fpath)
	check(err)

	file, err := os.Open(path)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elements []EncodedElement
	for scanner.Scan() {
		elements = append(elements, EncodedElement{Value: scanner.Text()})
	}
	return elements
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
