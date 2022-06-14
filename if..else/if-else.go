package main

import "fmt"

func average(a, b, c int) float32 {
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("quiz1 is higher than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz1 is lower than quiz2")
	} else {
		fmt.Println("quiz1 and quiz2 have the same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("acceptable grades")
	} else {
		fmt.Println("oh come onn")
	}
}
