package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("something failed while getting working directory")
	}

	file, err := os.Open(dir + "/defer_example.txt")
	if err != nil {
		log.Fatal("something failed while opening file.. %v", err)
	}

	defer file.Close()

	buffer := make([]byte, 50)
	bytes, err := file.Read(buffer)
	if err != nil {
		log.Fatal("something failed while reading file.. %v", err)
	}

	fmt.Printf("read number of bytes = %d\n", bytes)

}
