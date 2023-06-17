package main

import "fmt"

func main() {
	var word string
	var output string
	fmt.Println("Enter a Word")
	fmt.Scanln(&word)
	//inverts the word from right to left//
	for _, r := range word {
		output = string(r) + output
	}
	fmt.Println(output)

}
