//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func c_analyzWord(input *string, c chan int) {
	fmt.Printf("analyzing %v...\n", *input)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	length := len(*input)
	c <- length
}

func RunChannel() {
	fmt.Println("enter 3 words at least")

	channel := make(chan int, 3)

	var input1, input2, input3 string
	n, err := fmt.Scan(&input1, &input2, &input3)
	if err != nil {
		log.Fatal("something failed while reading input from user", err)
	}

	go c_analyzWord(&input1, channel)
	go c_analyzWord(&input2, channel)
	go c_analyzWord(&input3, channel)

	wordCount := 0
	letterCount := 0
	for {
		count := <-channel
		letterCount += count
		wordCount++

		if wordCount == n {
			break
		}
	}

	fmt.Printf("Word count: %d, total letter: %d", wordCount, letterCount)

}
