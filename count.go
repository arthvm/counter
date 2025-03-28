package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Counts struct {
	Bytes int
	Words int
	Lines int
}

func (c Counts) Print(w io.Writer, filename string) {
	fmt.Fprintf(w, "%d %d %d", c.Lines, c.Words, c.Bytes)

	if filename != "" {
		fmt.Fprintf(w, " %s", filename)
	}

	fmt.Fprintf(w, "\n")
}

func GetCounts(f io.ReadSeeker) Counts {
	const offsetStart = 0

	lineCount := CountLines(f)
	f.Seek(offsetStart, io.SeekStart)

	wordCount := CountWords(f)
	f.Seek(offsetStart, io.SeekStart)

	byteCount := CountBytes(f)

	return Counts{
		Bytes: byteCount,
		Words: wordCount,
		Lines: lineCount,
	}
}

func CountFile(filename string) (Counts, error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return Counts{}, err
	}

	counts := GetCounts(file)

	return Counts{
		Bytes: counts.Bytes,
		Words: counts.Words,
		Lines: counts.Lines,
	}, nil
}

func CountWords(file io.Reader) int {
	wordCount := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}

func CountLines(r io.Reader) int {
	lineCount := 0

	reader := bufio.NewReader(r)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if r == '\n' {
			lineCount++
		}
	}

	return lineCount
}

func CountBytes(r io.Reader) int {
	byteCount, _ := io.Copy(io.Discard, r)

	return int(byteCount)
}
