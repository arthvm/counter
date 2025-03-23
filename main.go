package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "./noexists.txt"

	log.SetFlags(0)

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	wordCount := CountWords(data)

	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)

	return len(words)
}
