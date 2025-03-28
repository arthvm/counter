package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/arthvm/counter/display"
)

type Counts struct {
	bytes int
	words int
	lines int
}

// Add will modify the values of the count by
// adding the values from the other
func (c Counts) Add(other Counts) Counts {
	c.bytes += other.bytes
	c.words += other.words
	c.lines += other.lines
	return c
}

func (c Counts) Print(w io.Writer, opts display.Options, suffixes ...string) {
	header := []string{}
	stats := []string{}

	if opts.ShouldShowLines() {
		header = append(header, "lines")
		stats = append(stats, strconv.Itoa(c.lines))
	}

	if opts.ShouldShowWords() {
		header = append(header, "words")
		stats = append(stats, strconv.Itoa(c.words))
	}

	if opts.ShouldShowBytes() {
		header = append(header, "bytes")
		stats = append(stats, strconv.Itoa(c.bytes))
	}

	hline := strings.Join(header, "\t") + "\t\n"
	sline := strings.Join(stats, "\t") + "\t"

	if opts.ShouldShowHeader() {
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

		res.bytes += size

		if r == '\n' {
			res.lines++
		}

		isSpace := unicode.IsSpace(r)

		if !isSpace && !isInsideWord {
			res.words++
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
		bytes: counts.bytes,
		words: counts.words,
		lines: counts.lines,
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
