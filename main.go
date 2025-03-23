package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := CountWords(data)

	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	wordCount := 0

	for _, byte := range data {
		if byte == ' ' {
			wordCount++
		}
	}

	wordCount++

	return wordCount
}
