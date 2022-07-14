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
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func sumFile(rd bufio.Reader) int {
	sum := 0
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			return sum
		}

		if err != nil {
			log.Fatal("something failed while reading line", err)
		}

		num, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Println("something failed while converting to string", err)
		}

		sum += num
	}
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	sum := 0

	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])
		if err != nil {
			log.Fatal("something failed while opening file", err)
		}

		rd := bufio.NewReader(file)

		calculate := func() {
			fileSum := sumFile(*rd)
			sum += fileSum
		}

		go calculate()
	}

	time.Sleep(300 * time.Millisecond)
	fmt.Println(sum)

}
