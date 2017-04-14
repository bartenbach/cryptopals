package main

import (
    "fmt"
    "path/filepath"
    "os"
    "bufio"
    "com/blakebartenbach/cryptopals/challenge-1"
    "strings"
    "github.com/hashicorp/vault/helper/xor"
)

func main() {
    // turn relative file path into absolute file path
    path, err := filepath.Abs("challenge-4/4.txt")
    check(err)

    // open the file
    file, err := os.Open(path)
    check(err)

    // defer closing file
//    defer file.Close()

    // create a file scanner to read lines

    // create alphabet array
    alphabet := [68]string{ "A","B","C","D","E","F","G","H","I","J","K","L",
            "M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","0","1",
            "2","3","4","5","6","7","8","9","~","!","@","#","$","%","^","&",
            "*","(",")","-","_","=","+","{","}","[","]","\\","|","/",".","<",
            ">",",","`","?","\"","\"",":",";"}

    // loop through the entire alphabet testing
    for i:=0; i<68; i++ {
        scanner := bufio.NewScanner(file)
        fmt.Println(i)
        var letter string = alphabet[i]
        checkAgainstFile(&letter,scanner)
    }
}

func checkAgainstFile(letter *string, scanner *bufio.Scanner) {
    // create array to XOR bytes against
    s0 := strings.Repeat(*letter, 30)
    lowercase := strings.ToLower(*letter)
    lowercasebyte := string(lowercase)
    s1 := strings.Repeat(lowercasebyte, 30)
    // scan file line by line and XOR against array
    for scanner.Scan() {
        var decoded = challenge_1.DecodeHex(scanner.Text())
        //var xored = challenge_2.XORvalues(s0, decoded)
        fmt.Println(s0)
        fmt.Println(s1)
        fmt.Println(string(decoded))
        var xored, _ = xor.XORBase64(s0, string(decoded))
        //var xoredLower = challenge_2.XORvalues(s1, decoded)
        var xoredLower, _ = xor.XORBase64(s1,string(decoded))
        var xstring string = string(xored)
        var xstringLower string = string(xoredLower)

        // attempt to cut down on crap by scoring the english language
        // likely any english sentence will contain an "e" whereas most
        // of the junk that comes out does not...
        //if (strings.Contains(xstring, "e")) {
            fmt.Println(xstring)
        //}
        //if (strings.Contains(xstringLower, "e")) {
            fmt.Println(xstringLower)
       // }
    }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
