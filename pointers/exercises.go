//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type SecurityTag bool

type Product struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = active
}

func deactivate(tag *SecurityTag) {
	*tag = inactive
}

func checkout(products []Product) {
	fmt.Println("checking out...")
	// use index style for loop instead of range style, range copies the values
	for i := 0; i < len(products); i++ {
		deactivate(&products[i].tag)
	}
}

const (
	active   = true
	inactive = false
)

func main() {

	item1 := Product{"Iphone 10", active}
	item2 := Product{"Iphone 11", active}
	item3 := Product{"Iphone 12", active}
	item4 := Product{"Iphone 13", active}

	items := []Product{item1, item2, item3, item4}
	fmt.Println(items)

	deactivate(&items[3].tag)
	fmt.Println(items)

	checkout(items)
	fmt.Println(items)
}
