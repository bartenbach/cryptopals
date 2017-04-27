package challenge_6

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

var table = [256]uint8{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

type EncodedBlock struct {
	Block []byte
}

func Challenge6() {
	// Read in elements from file and decode them
	var elements []byte = GetStringFromFile("challenge-6/6.txt")

	// set variables to remember the shortest hamming distance, and the corresponding keysize
	shortestDistance := 999
	shortestKeysize := 0

	// 'i' represents the guessed keysizes, guessing from 4 to 40
	for i := 4; i <= 40; i++ {
		// divide by (keysize) to normalize the result
		var distance int = CalculateEditDistance(i, elements) / i
		if distance < shortestDistance {
			fmt.Printf("Shortest distance: %d\n", shortestDistance)
			fmt.Printf("New shortest distance: %d\n", distance)
			shortestDistance = distance
			shortestKeysize = i
		}
	}

	fmt.Printf("The shortest hamming distance found was: %d\n", shortestDistance)
	fmt.Printf("The corresponding keysize was: %d\n", shortestKeysize)

	fmt.Printf("\n\nBreaking up ciphertext into <keysize> blocks...\n\n")

	var blocks []EncodedBlock = GetEncodedBlocks(elements)
}

func GetEncodedBlocks(elements []byte) []EncodedBlock {

}

func CalculateEditDistance(keysize int, elements []byte) int {
	bytes1 := elements[0:keysize]
	bytes2 := elements[keysize : keysize*2]
	fmt.Printf("keysize worth of bytes1: ")
	PrintPrettySlice(bytes1)
	fmt.Printf("keysize worth of bytes2: ")
	PrintPrettySlice(bytes2)
	hammingd, err := HammingDistance(bytes1, bytes2)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return 9999
	} else {
		fmt.Printf("Hamming Distance: %d\n", hammingd)
		return hammingd
	}
}

func HammingDistance(x, y []byte) (int, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("String lengths not equal! %d != %d", len(x), len(y))
	}

	var hammingd int = 0

	for i := range x {
		hammingd += int(table[(x[i] ^ y[i])])
	}
	return hammingd, nil
}

// Uses printf to print evenly spaced columns for slices
func PrintPrettySlice(b []byte) {
	fmt.Printf("[")
	for x := range b {
		fmt.Printf("%4d", b[x])
	}
	fmt.Printf(" ]\n")
}

// reads elements from file and appends them to a list of bytes
// it also decodes them from base64 because that's convenient
func GetStringFromFile(fpath string) []byte {
	path, err := filepath.Abs(fpath)
	check(err)

	file, err := os.Open(path)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elements []byte
	for scanner.Scan() {
		element, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			fmt.Printf("ERROR: %s", err)
		} else {
			elements = append(elements, []byte(element)...)
		}
	}
	return elements
}

// error boilerplate
func check(e error) {
	if e != nil {
		panic(e)
	}
}
