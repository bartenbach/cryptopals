package challenge_6

import (
	"bufio"
	"bytes"
	"com/blakebartenbach/cryptopals/challenge-2"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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

var b64 = [38]byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
	'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1',
	'2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
}

type EncodedBlock struct {
	Block []byte
}

// TODO: once solved, this has a TON of debug output that should be removed
func Challenge6() {
	// Read in elements from file and decode them
	var elements []byte = GetStringFromFile("challenge-6/6.txt")

	// set variables to remember the shortest hamming distance, and the corresponding keysize
	shortestDistance := 999.0
	shortestKeysize := 0
	var keysizes []int

	// 'i' represents the guessed keysizes, guessing from 2 to 40
	for i := 2; i <= 40; i++ {
		// divide by (keysize) to normalize the result
		var distance float64 = CalculateEditDistance(i, elements) / float64(i)
		keysizes = append(keysizes, i)
		if distance < shortestDistance {
			fmt.Printf("Shortest distance: %f\n", shortestDistance)
			fmt.Printf("New shortest distance: %f\n", distance)
			shortestDistance = distance
			shortestKeysize = i
		}
	}

	sort.Ints(keysizes)
	fmt.Println("SORTED KEYSIZES")
	fmt.Println(keysizes)

	fmt.Printf("\n\nThe shortest hamming distance found was: %f\n", shortestDistance)
	fmt.Printf("The corresponding keysize was: %d\n", shortestKeysize)

	fmt.Printf("\n\nBreaking up ciphertext into <keysize> blocks...\n\n")

	// split up ciphertext into <keysize> length blocks
	var blocks []EncodedBlock = GetEncodedBlocks(shortestKeysize, elements)
	fmt.Printf("\n\nLENGTH OF BLOCKS: %d\n\n", len(blocks))

	for i := 0; i < 5; i++ {
		var x []byte
		for _, block := range blocks {
			x = append(x, block.Block[i])
		}
		SingleCharacterXOR(x)
		fmt.Printf("%s\n", string(x))
	}
}

func SingleCharacterXOR(x []byte) {
	for i := range b64 {
		// get our current character
		var letter = b64[i]
		fmt.Println("Current letter " + string(letter))

		s0 := bytes.Repeat([]byte{letter}, len(x))
		var xored, err = challenge_2.XORvalues(x, s0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(xored))
	}
}

func GetEncodedBlocks(keysize int, elements []byte) []EncodedBlock {
	var blocks []EncodedBlock
	for j := 0; j < len(elements)/keysize; j += keysize {
		bytes1 := elements[j : j+keysize]
		blocks = append(blocks, EncodedBlock{Block: bytes1})
	}
	return blocks
}

func CalculateEditDistance(keysize int, elements []byte) float64 {
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
		fmt.Printf("Hamming Distance: %f\n", hammingd)
		return hammingd
	}
}

func HammingDistance(x, y []byte) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf("String lengths not equal! %d != %d", len(x), len(y))
	}

	var hammingd float64 = 0

	for i := range x {
		hammingd += float64(table[(x[i] ^ y[i])])
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
		// do i decode from base64 here or later?  does it matter?
		//err = nil
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
