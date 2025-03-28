package main_test

import (
	"bytes"
	"strings"
	"testing"

	counter "github.com/arthvm/counter"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "new lines",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "multiple spaces",
			input: "This is a sentence.  This is another",
			wants: 7,
		},
		{
			name:  "prefixed multiple spaces",
			input: "  Hello",
			wants: 1,
		},
		{
			name:  "suffixed multiple spaces",
			input: "Hello      ",
			wants: 1,
		},
		{
			name:  "tab character",
			input: "Hello\tWorld\n",
			wants: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			res := counter.GetCounts(r).Words

			if res != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "simple 5 words, 1 new line",
			input: "one two three four five\n",
			wants: 1,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "no new lines",
			input: "one two three four five",
			wants: 0,
		},
		{
			name:  "no new line at end",
			input: "one two three four five\nsix",
			wants: 1,
		},
		{
			name:  "multi newline string",
			input: "\n\n\n\n",
			wants: 4,
		},
		{
			name:  "multi word line string",
			input: "1\n2\n3\n4\n5\n",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			res := counter.GetCounts(r).Lines
			if res != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "five words",
			input: "one two three four five",
			wants: 23,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "all spaces",
			input: "       ",
			wants: 7,
		},
		{
			name:  "newlines tabs and words",
			input: "one\ntwo\nthree\nfour\t\n",
			wants: 20,
		},
		{
			name:  "unicode characters",
			input: "ẂѶ",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			res := counter.GetCounts(r).Bytes
			if res != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestGetCounts(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants counter.Counts
	}{
		{
			name:  "five words",
			input: "one two three four five\n",
			wants: counter.Counts{
				Lines: 1,
				Words: 5,
				Bytes: 24,
			},
		},
		{
			name:  "five words no new line",
			input: "one two three four five",
			wants: counter.Counts{
				Lines: 0,
				Words: 5,
				Bytes: 23,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			res := counter.GetCounts(r)
			if res != tc.wants {
				t.Logf("expected: %v got: %v", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestPrintCounts(t *testing.T) {
	type inputs struct {
		counts    counter.Counts
		filenames []string
	}

	testCases := []struct {
		name  string
		input inputs
		wants string
	}{
		{
			name: "simple five words.txt",
			input: inputs{
				counts: counter.Counts{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				filenames: []string{"words.txt"},
			}, wants: "1 5 24 words.txt\n",
		},
		{
			name: "no filename",
			input: inputs{
				counts: counter.Counts{
					Lines: 1,
					Words: 4,
					Bytes: 20,
				},
			}, wants: "1 4 20\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			tc.input.counts.Print(buffer, tc.input.filenames...)

			if buffer.String() != tc.wants {
				t.Logf("expected: %v got: %v", []byte(tc.wants), buffer.Bytes())
				t.Fail()
			}
		})
	}
}

func TestAddCounts(t *testing.T) {
	type inputs struct {
		counts counter.Counts
		other  counter.Counts
	}

	testCases := []struct {
		name  string
		input inputs
		wants counter.Counts
	}{
		{
			name: "simple add by one",
			input: inputs{
				counts: counter.Counts{
					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				other: counter.Counts{
					Lines: 1,
					Words: 1,
					Bytes: 1,
				},
			},
			wants: counter.Counts{
				Lines: 2,
				Words: 6,
				Bytes: 25,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			totals := tc.input.counts
			res := totals.Add(tc.input.other)

			if res != tc.wants {
				t.Logf("expected: %v got: %v", tc.wants, totals)
				t.Fail()
			}
		})
	}
}
