package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	header := []string{}
	stats := []string{}

	if opts.ShouldShowLines() {
		header = append(header, "lines")
		stats = append(stats, strconv.Itoa(c.Lines))
	}

	if opts.ShouldShowWords() {
		header = append(header, "words")
		stats = append(stats, strconv.Itoa(c.Words))
	}

	if opts.ShouldShowBytes() {
		header = append(header, "bytes")
		stats = append(stats, strconv.Itoa(c.Bytes))
	}

	hline := strings.Join(header, "\t") + "\t\n"
	sline := strings.Join(stats, "\t") + "\t"

	if opts.ShowHeader {
		fmt.Fprint(w, hline)
	}
	fmt.Fprint(w, sline)

	suffixStr := strings.Join(suffixes, " ")
	if suffixStr != "" {
		fmt.Fprintf(w, " %s", suffixStr)
	}

	fmt.Fprintf(w, "\n")
}

func GetCounts(f io.Reader) Counts {
	res := Counts{}

	isInsideWord := false
	reader := bufio.NewReader(f)

	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			break
		}

		res.Bytes += size

		if r == '\n' {
			res.Lines++
		}

		isSpace := unicode.IsSpace(r)

		if !isSpace && !isInsideWord {
			res.Words++
		}

		isInsideWord = !isSpace
	}

	return res
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
