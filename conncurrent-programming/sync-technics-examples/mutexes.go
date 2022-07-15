package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type total struct {
	words   []string
	letters int
	sync.Mutex
}

func analyzWord(wg *sync.WaitGroup, input *total, iteration int) {
	// waiting
	sleep()
	word := input.words[iteration]
	fmt.Printf("analyzing %v...\n", word)
	input.Lock()
	defer input.Unlock()
	defer wg.Done()

	input.letters += len(word)
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func RunMutexes() {
	rand.Seed(time.Now().UnixNano())

	words := total{}

	var wg sync.WaitGroup
	var input1, input2, input3 string

	fmt.Println("enter at least 3 words")
	n, err := fmt.Scan(&input1, &input2, &input3)
	if err != nil {
		log.Fatal("something failed while reading inputs", err)
	}

	words.words = append(words.words, input1, input2, input3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go analyzWord(&wg, &words, i)
	}

	wg.Wait()
	fmt.Printf("total words: %d, letters count: %d", n, words.letters)

}
