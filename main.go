package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := 1

	for _, byte := range data {
		if byte == ' ' {
			wordCount++
		}
	}

	fmt.Println(wordCount)
}
