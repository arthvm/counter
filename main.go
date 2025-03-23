package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "./lots_of_words.txt"

	log.SetFlags(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	PrintFileContents(file)

	// wordCount := CountWords(data)
	//
	// fmt.Println(wordCount)
}

func PrintFileContents(file *os.File) {
	const bufferSize = 4096 * 2
	buffer := make([]byte, bufferSize)

	totalSize := 0
	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		totalSize += size
	}

	fmt.Println("total bytes read:", totalSize)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)

	return len(words)
}
