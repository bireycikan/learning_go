package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index uint) (int, error) {
	if int(index) > len(s.values) {
		// we can use errors.New for raw string, but if we use with fmt.Sprintf function, it complains about the usage of errors.New
		// return 0, errors.New(fmt.Sprintf("no element at index %v", index))

		// instead, we can use directly fmt.Errorf() function on the other hand
		return 0, fmt.Errorf("no element at index %v", index)
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{}
	stuff.values = []int{1, 2, 3, 4, 5, 6}

	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("value is: %d", value)
}
