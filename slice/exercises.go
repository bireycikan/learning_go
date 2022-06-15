//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printContents(line []Part) {

	fmt.Println()
	for i := 0; i < len(line); i++ {
		item := line[i]
		fmt.Println(item)
	}
}

func main() {
	line := []Part{"part1", "part2", "part3"}
	printContents(line)

	// other technic
	otherLine := make([]Part, 3)
	otherLine[0] = "other part 1"
	otherLine[1] = "other part 2"
	otherLine[2] = "other part 3"

	line = append(line, "part4", "part5")
	printContents(line)

	line = line[3:]
	printContents(line)

	// other print
	printContents(otherLine)

}
