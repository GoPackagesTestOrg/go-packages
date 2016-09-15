package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World"

	// Create a new buffered scanner that scans from a
	// new strings reader that reads from string s
	ss := bufio.NewScanner(strings.NewReader(s))

	// Use bufio.ScanBytes as the scanners splitter
	// which will split the string into seperate bytes
	ss.Split(bufio.ScanBytes)

	// Scan the first token from ss
	ss.Scan()

	// Print out the token text received from ss.Scan()
	fmt.Println(ss.Text()) // H

	// Scan the second token from ss
	ss.Scan()

	// Print out the token text received from ss.Scan()
	fmt.Println(ss.Text()) // e

	// Print out the most recent token generated from scanning with ss
	//
	// Bytes returns the most recent token generated by a call to Scan.
	// The underlying array may point to data that will be overwritten
	// by a subsequent call to Scan. It does no allocation.
	fmt.Println(string(ss.Bytes())) // e

	// Check if there were any scanner errors
	if err := ss.Err(); err != nil {
		// If there was any erorrs then log them
		log.Fatalln(err)
	}
}