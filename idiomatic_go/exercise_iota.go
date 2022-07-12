//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation byte

const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
)

func (o Operation) calculate(num1, num2 int) int {
	switch o {
	case Add:
		return num1 + num2
	case Subtract:
		return num1 - num2
	case Multiply:
		return num1 * num2
	case Divide:
		return num1 / num2
	}

	panic("unhandled operation")
}

func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := Operation(Subtract)
	fmt.Println(sub.calculate(10, 3)) // = 7

	mul := Operation(Multiply)
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := Operation(Divide)
	fmt.Println(div.calculate(100, 2)) // = 50
}
