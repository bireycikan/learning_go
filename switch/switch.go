package main

import "fmt"

func price() int {
	return 50
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("moderately priced item")
	default:
		fmt.Println("expensive item")
	}

	ticket := Business

	switch ticket {
	case Economy:
		fmt.Println("Economy ticket")
	case Business:
		fmt.Println("Business ticket")
	case FirstClass:
		fmt.Println("First class ticket")
	default:
		fmt.Println("other seating")
	}

}
