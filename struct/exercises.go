//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

func width(rec Rectangle) int {
	return rec.b.x - rec.a.x
}

func height(rec Rectangle) int {
	return rec.a.y - rec.b.y
}

func calculateArea(rec Rectangle) int {
	return width(rec) * height(rec)
}

func calculatePerimeter(rec Rectangle) int {
	return 2 * (width(rec) + height(rec))
}

func main() {
	topLeft := Coordinate{2, 10}
	bottomRight := Coordinate{10, 5}
	rectangle := Rectangle{a: topLeft, b: bottomRight}
	// calculation
	area := calculateArea(rectangle)
	fmt.Println("Area is", area)
	perimeter := calculatePerimeter(rectangle)
	fmt.Println("Perimeter is", perimeter)

	rectangle.a.y *= 2
	rectangle.b.x *= 2
	area2 := calculateArea(rectangle)
	fmt.Println("New area is", area2)
	perimeter2 := calculatePerimeter(rectangle)
	fmt.Println("New perimeter is", perimeter2)
}
