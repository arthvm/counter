package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
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

func GetCounts(r io.Reader) Counts {
	bytesReader, bytesWriter := io.Pipe()
	wordsReader, wordsWriter := io.Pipe()
	linesReader, linesWriter := io.Pipe()

	w := io.MultiWriter(bytesWriter, wordsWriter, linesWriter)

	chBytes := make(chan int)
	chWords := make(chan int)
	chLines := make(chan int)

	go func() {
		defer close(chBytes)
		chBytes <- CountBytes(bytesReader)
	}()
	go func() {
		defer close(chWords)
		chWords <- CountWords(wordsReader)
	}()
	go func() {
		defer close(chLines)
		chLines <- CountLines(linesReader)
	}()

	io.Copy(w, r)
	bytesWriter.Close()
	wordsWriter.Close()
	linesWriter.Close()

	byteCount := <-chBytes
	wordCount := <-chWords
	lineCount := <-chLines

	return Counts{
		bytes: byteCount,
		words: wordCount,
		lines: lineCount,
	}
}

func GetCountsSinglePass(f io.Reader) Counts {
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

type FileCountsResult struct {
	Counts   Counts
	Filename string
	Err      error
	Idx      int
}

func CountFiles(filenames []string) <-chan FileCountsResult {
	ch := make(chan FileCountsResult)

	wg := sync.WaitGroup{}
	wg.Add(len(filenames))

	for i, filename := range filenames {
		go func() {
			defer wg.Done()
			res, err := CountFile(filename)

			ch <- FileCountsResult{
				Counts:   res,
				Filename: filename,
				Err:      err,
				Idx:      i,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
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
