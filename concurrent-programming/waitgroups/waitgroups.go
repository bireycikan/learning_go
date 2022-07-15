package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		counter += 1
		go func() {
			defer func() {
				fmt.Println("goroutines remaining", counter)
				counter -= 1
				wg.Done()
			}()

			duration := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("waiting for", duration)
			time.Sleep(duration)
		}()
	}

	fmt.Println("waiting for goroutines to finish")
	wg.Wait()
	fmt.Println("done!")
}
