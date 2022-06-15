package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {

	route := []string{"Grocery", "Department Store", "Salon"}
	printSlice("Route 1", route)

	route = append(route, "Home")
	printSlice("Route 2", route)

	anotherRoute := []string{"Library", "Laboratory"}
	route = append(route, anotherRoute...)
	printSlice("Route 3", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	route = route[2:]
	printSlice("Route 4", route)

	route = route[:3]
	printSlice("Route 5", route)

}
