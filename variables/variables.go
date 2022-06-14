package main

import "fmt"

func main() {
	// type inference
	var myName = "Birey"
	// inserts white space at the end of the first string argument
	fmt.Println("My name is", myName, myName)

	// type annotation
	var name string = "Jessica"
	fmt.Println("name =", name)

	// create and assign
	username := "Kate"
	fmt.Println("her name is", username)

	// create and uninitialize it
	var sum int
	fmt.Println("sum is", sum)

	// multiple assignment
	var a, b byte = 1, 2
	fmt.Println("a =", a, "b =", b)

	var (
		x bool   = false
		y string = "let's go"
		z int    = 5
	)
	fmt.Println("x is", x, "y is", y, "z is", z)

	// multiple create and assign
	c, d := false, "true"
	fmt.Println("c =", c, "d =", d)

	// re-assign d variable, (special case for second variable with := and comma , operators)
	e, d := true, "false"
	fmt.Println("e =", e, "d =", d)

	// re-assign sum
	sum = z + 1
	fmt.Println("sum is", sum)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
