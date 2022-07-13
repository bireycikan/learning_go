//--Summary:
//  Create a program that can create a report of rune information from

////  lines of text.
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type printInfo func(line string)

func report(lines []string, callback printInfo) {
	for i := 0; i < len(lines); i++ {
		callback(lines[i])
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	letterCount := 0
	digitCount := 0
	spaceCount := 0
	puncCount := 0

	lineCallback := func(line string) {
		for _, str := range line {
			if unicode.IsLetter(str) {
				letterCount++
			} else if unicode.IsDigit(str) {
				digitCount++
			} else if unicode.IsSpace(str) {
				spaceCount++
			} else if unicode.IsPunct(str) {
				puncCount++
			}
		}
	}

	report(lines, lineCallback)

	fmt.Printf("letterCount: %d, digitCount: %d, spaceCount: %d, puncCount: %d", letterCount, digitCount, spaceCount, puncCount)
}
