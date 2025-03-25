package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	filename := "./words.txt"

	log.SetFlags(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	wordCount := CountWordsInFile(file)
	fmt.Println(wordCount)
}

func CountWordsInFile(file *os.File) int {
	wordCount := 0
	isInsideWord := false

	const bufferSize = 4096 * 2
	buffer := make([]byte, bufferSize)

	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		isInsideWord = !unicode.IsSpace(rune(buffer[0])) && isInsideWord

		bufferCount := CountWords(buffer[:size])
		if isInsideWord {
			bufferCount -= 1
		}

		wordCount += bufferCount

		isInsideWord = !unicode.IsSpace(rune(buffer[size-1]))
	}

	return wordCount
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)

	return len(words)
}
