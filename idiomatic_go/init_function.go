package main

import (
	"fmt"
	"regexp"
)

var EmailExpr *regexp.Regexp

// all packages will execute init() function before main() runs
// each package can have its own init() function
func init() {
	compiled, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("failed to compile regular expression")
	}

	EmailExpr = compiled
	fmt.Println("regular expression compiled successfully")
}

func main() {
	fmt.Println(isValidEmail("invalid"))
	fmt.Println(isValidEmail("valid@example.com"))
	fmt.Println(isValidEmail("invalid@example"))
}

func isValidEmail(addr string) bool {
	return EmailExpr.Match([]byte(addr))
}
