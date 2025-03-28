package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Counts struct {
	Bytes int
	Words int
	Lines int
}

// Add will modify the values of the count by
// adding the values from the other
func (c Counts) Add(other Counts) Counts {
	c.Bytes += other.Bytes
	c.Words += other.Words
	c.Lines += other.Lines
	return c
}

func (c Counts) Print(w io.Writer, opts DisplayOptions, suffixes ...string) {
	xs := []string{}

	if opts.ShouldShowLines() {
		xs = append(xs, strconv.Itoa(c.Lines))
	}

	if opts.ShouldShowWords() {
		xs = append(xs, strconv.Itoa(c.Words))
	}

	if opts.ShouldShowBytes() {
		xs = append(xs, strconv.Itoa(c.Bytes))
	}

	xs = append(xs, suffixes...)

	line := strings.Join(xs, " ")

	fmt.Fprintln(w, line)
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
