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

	wordDetected := false
	wordCount := 0

	for _, byte := range data {
		if byte == ' ' {
			wordCount++
		} else {
			wordDetected = true
		}
	}

	if !wordDetected {
		return 0
	}

	wordCount++

	return wordCount
}
