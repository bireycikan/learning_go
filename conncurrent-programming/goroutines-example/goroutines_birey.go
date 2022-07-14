//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func getTotal(filename string, total *int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println("something failed while converting to string", err)
		}
		*total += num
	}

	if err := sc.Err(); err != nil {
		log.Fatal("scan failed", err)
	}

}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	total := 0
	for _, file := range files {
		go getTotal(file, &total)
	}

	time.Sleep(1 * time.Second)

	fmt.Println(total)

}
