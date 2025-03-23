package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	fmt.Println("data:", string(data))
}
