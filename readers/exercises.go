//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	CmdHello = "hello"
	CmdBye   = "bye"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	numNonBlank := 0
	numCommands := 0

	for {
		input, inputErr := r.ReadString('\n')
		n := strings.TrimSpace(input)

		if n == "" {
			continue
		} else {
			numNonBlank++
		}

		if n == "Q" || n == "q" {
			break
		} else {
			switch n {
			case CmdHello:
				fmt.Println("command response: hi")
				numCommands++
			case CmdBye:
				fmt.Println("command response: bye")
				numCommands++
			}
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Printf("Something failed reading inputs: %v", inputErr)
		}
	}

	fmt.Printf("You entered: %d lines,\nYou entered: %d commands", numNonBlank, numCommands)
}
