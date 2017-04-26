package challenge_6

import (
//	"encoding/base64"
	"fmt"
    "path/filepath"
    "os"
    "bufio"
    "encoding/base64"
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

func Challenge6() {
	// Read in elements from encrypted file - store them in a slice of EncodedElements (really just strings)
    var elements []byte = GetStringFromFile("challenge-6/6.txt")
    PrintPrettySlice(elements)

	//Guess our keysize.  Will have to find the keysize with the shortest hamming distance.
	// should I do this manually or programmatically?
	keysize := 4

	// Using our keysize, calculate hamming distance between two keys
    fmt.Printf("LENGTH OF STRING: %d\n", len(elements))
    for j:=0; j <= (len(elements)/keysize); j+=(keysize*2) {
        bytes1 := elements[j : j+keysize]
        bytes2 := elements[j+keysize : j+(keysize*2)]
        fmt.Printf("keysize worth of bytes1: ")
        PrintPrettySlice(bytes1)
        fmt.Printf("keysize worth of bytes2: ")
        PrintPrettySlice(bytes2)
        hammingd, err := HammingDistance(bytes1,bytes2)
        if err != nil {
            fmt.Printf("ERROR: %s", err)
        } else {
            fmt.Printf("Hamming Distance: %d\n", hammingd)
        }
     }
        // need to get the first <keysize> worth of bytes from the element
		// wow...this is confusing
//		bytes1 := elements[j : j+keysize]
//		bytes2 := elements[j+keysize : j+(keysize*2)]
//		hammingd, err := HammingDistance(bytes1, bytes2)
//		if err != nil {
//			println("ERROR: ", err)
//		} else {
//			// normalize hamming distance by dividing by the keysize
//			fmt.Printf("Hamming distance: %d\n", hammingd/keysize)
//		}

//	}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
