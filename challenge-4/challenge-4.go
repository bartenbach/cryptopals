package main

import (
    "fmt"
    "path/filepath"
    "os"
    "bufio"
    "com/blakebartenbach/cryptopals/challenge-1"
    "bytes"
    "com/blakebartenbach/cryptopals/challenge-2"
    "strings"
)

func main() {
    // turn relative file path into absolute file path
    path, err := filepath.Abs("challenge-4/4.txt")
    check(err)

    // open the file
    file, err := os.Open(path)
    check(err)

    // defer closing file
    defer file.Close()

    // create a file scanner to read lines
    scanner := bufio.NewScanner(file)

    // create alphabet array
    alphabet := [26]byte{ 'A','B','C','D','E','F','G','H','I','J','K','L',
            'M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'}

    // loop through the entire alphabet testing
    for i:=0; i<26; i++ {
        letter := byte(alphabet[i])
        checkAgainstFile(letter,scanner)

    }
}

func checkAgainstFile(letter byte, scanner *bufio.Scanner) {
    // create array to XOR bytes against
        s0 := bytes.Repeat([]byte{letter}, 30)
        lowercase := strings.ToLower(string(letter))
        lowercasebyte := []byte(lowercase)
        s1 := bytes.Repeat([]byte{lowercasebyte[0]}, 30)
        // scan file line by line and XOR against array
        for scanner.Scan() {
            var decoded = challenge_1.DecodeHex(scanner.Text())
            var xored = challenge_2.XORvalues(s0, decoded)
            var xoredLower = challenge_2.XORvalues(s1, decoded)
            var xstring string = string(xored)
            var xstringLower string = string(xoredLower)

            // attempt to cut down on crap by scoring the english language
            // likely any english sentence will contain an 'e' whereas most
            // of the junk that comes out does not...
            //if (strings.Contains(xstring, "e")) {
                fmt.Println(xstring)
            //}
            //if (strings.Contains(xstringLower, "e")) {
                fmt.Println(xstringLower)
            //}
        }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
