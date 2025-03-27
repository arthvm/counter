package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	total := 0
	filenames := os.Args[1:]
	didError := false

	for _, filename := range filenames {
		counts, err := CountFile(filename)
		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}

		total += counts.Words

		fmt.Println(counts.Lines, counts.Words, counts.Bytes, filename)
	}

	if len(filenames) == 0 {
		counts := GetCounts(os.Stdin)
		fmt.Println(counts.Lines, counts.Words, counts.Bytes)
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

	if didError {
		os.Exit(1)
	}
}
