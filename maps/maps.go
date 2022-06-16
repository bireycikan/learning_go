package main

import (
	"fmt"
)

func main() {

	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["break"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk delete, new list is", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")

	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yeah", cereal, "boxes")
	}

	totalItems := 0

	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("There are", totalItems, "in the shopping list")

}
