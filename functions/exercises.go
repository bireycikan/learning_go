//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hello", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func getSameMessage(message string) string {
	return message
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func add(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func returnOneNumber() int8 {
	return 2
}

//* Write a function that returns any two numbers
func returnTwoNumbers() (uint16, uint32) {
	return 34225, 3434534523
}

func main() {

	// greet
	greetPerson("birey")

	// same message
	fmt.Println("secret message is", getSameMessage("hello world"))

	// add numbers
	total := add(1, 2, 3)
	fmt.Println("total is", total)

	// return any number
	var anyNumber = returnOneNumber()
	fmt.Println("any number", anyNumber)

	// return any two numbers
	anyNum1, anyNum2 := returnTwoNumbers()
	fmt.Println("anyNum1", anyNum1, "anyNum2", anyNum2)

	//* Add three numbers together using any combination of the existing functions.
	combination := add(total, int(anyNum1), int(anyNum2))
	//  * Print the result
	fmt.Println("combination is", combination)
	//* Call every function at least once
}
