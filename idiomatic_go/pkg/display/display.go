package display

import "fmt"

// if we name our function with a capital letter, it means we can access this function from other packages
func Display(msg string) {
	fmt.Println(msg)
}

// if we name our function with a lower case letter, it means we cannot access this function from other packages
func hello(msg string) {
	fmt.Println(msg)
}
