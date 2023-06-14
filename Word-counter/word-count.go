package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func main() {

	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 3
	res := count(b)

	if res != exp {
		fmt.Println("expected 3 word got more")
	}

}

func count(r io.Reader) int {

	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc

}
