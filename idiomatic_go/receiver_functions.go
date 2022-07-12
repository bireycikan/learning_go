package main

import "fmt"

type Coordinate struct {
	X, Y int
}

// normal function
func shifyBy(x, y int, coord *Coordinate) Coordinate {
	coord.X += x
	coord.Y += y

	return *coord
}

// pointer receiver function
func (coord *Coordinate) shiftBy(x, y int) Coordinate {
	coord.X += x
	coord.Y += y

	return *coord
}

// value receiver function
func (c Coordinate) Dist(other Coordinate) Coordinate {
	return Coordinate{c.X - other.X, c.Y - other.Y}
}

func main() {
	coord := Coordinate{5, 5}
	shifted := shifyBy(1, 1, &coord)

	fmt.Println(shifted)

	// for receiver
	shifted_receiver := coord.shiftBy(1, 1)

	fmt.Println(shifted_receiver)

	first := Coordinate{2, 2}
	second := Coordinate{1, 5}
	distance := first.Dist(second)

	fmt.Println(distance)
}
