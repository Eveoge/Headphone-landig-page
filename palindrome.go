package main

import (
	"fmt"
	"strings"
)

func main() {

	var word string
	var output string

	fmt.Println("Enter a String")

	fmt.Scanln(&word)
	fmt.Println("Checking if word is palindrome")

	word = strings.TrimSpace(word)

	fmt.Println("word")
	for _, r := range word {
		output = string(r) + output
	}

	if word == output {
		fmt.Println(word, "is palindrome")
	} else {
		fmt.Println(word, "is not palindrome")
	}
}
