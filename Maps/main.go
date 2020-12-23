package main

import "fmt"

func main() {

	// Make a map and assign value
	numbers := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(numbers)

	// Make a map
	colours := make(map[string]string)

	// Assign a value
	colours["red"] = "#ff0000"
	fmt.Println(colours)

	// Delete a value
	delete(colours, "red")
	fmt.Println(colours)

	// Iterating throgh map
	printMap(numbers)
}

func printMap(n map[string]int) {
	for name, value := range n {
		fmt.Printf("Value of text %v is %v\n", name, value)
	}
}
