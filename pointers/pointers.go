package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	// dereferencing is made automatically for struct
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	// if the type is the same what we send, we do not need to make any special thing by passing into the function
	increment(counter)
}

func main() {

	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Print(phrase)
}
