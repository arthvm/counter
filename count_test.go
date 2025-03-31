package counter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/arthvm/counter/display"
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
			res := GetCounts(r).words

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

			res := GetCounts(r).lines
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

			res := GetCounts(r).bytes
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
		wants Counts
	}{
		{
			name:  "five words",
			input: "one two three four five\n",
			wants: Counts{
				lines: 1,
				words: 5,
				bytes: 24,
			},
		},
		{
			name:  "five words no new line",
			input: "one two three four five",
			wants: Counts{
				lines: 0,
				words: 5,
				bytes: 23,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			res := GetCounts(r)
			if res != tc.wants {
				t.Logf("expected: %v got: %v", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestPrintCounts(t *testing.T) {
	type inputs struct {
		counts    Counts
		opts      display.NewOptionsArgs
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
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: true,
					ShowWords: true,
					ShowLines: true,
				},
				filenames: []string{"words.txt"},
			},
			wants: "1\t5\t24\t words.txt\n",
		},
		{
			name: "no filename",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 4,
					bytes: 20,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: true,
					ShowWords: true,
					ShowLines: true,
				},
			},
			wants: "1\t4\t20\t\n",
		},
		{
			name: "simple five words.txt no option",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				filenames: []string{"words.txt"},
			},
			wants: "1\t5\t24\t words.txt\n",
		},
		{
			name: "simple five words.txt show lines",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: false,
					ShowWords: false,
					ShowLines: true,
				},
				filenames: []string{"words.txt"},
			},
			wants: "1\t words.txt\n",
		},
		{
			name: "simple five words.txt show words",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: false,
					ShowWords: true,
					ShowLines: false,
				},
				filenames: []string{"words.txt"},
			},
			wants: "5\t words.txt\n",
		},
		{
			name: "simple five words.txt show bytes",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				opts: display.NewOptionsArgs{
					ShowBytes: true,
					ShowWords: false,
					ShowLines: false,
				},
				filenames: []string{"words.txt"},
			},
			wants: "24\t words.txt\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			tc.input.counts.Print(buffer, display.NewOptions(tc.input.opts), tc.input.filenames...)

			if buffer.String() != tc.wants {
				t.Logf("expected: %v got: %v", tc.wants, buffer.String())
				t.Fail()
			}
		})
	}
}

func TestAddCounts(t *testing.T) {
	type inputs struct {
		counts Counts
		other  Counts
	}

	testCases := []struct {
		name  string
		input inputs
		wants Counts
	}{
		{
			name: "simple add by one",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				other: Counts{
					lines: 1,
					words: 1,
					bytes: 1,
				},
			},
			wants: Counts{
				lines: 2,
				words: 6,
				bytes: 25,
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

var benchData = []string{
	"This is a test data string\nthat spans across\nmultiple lines\n",
	"one two three\nfour five\nsix\nseven\neight\n",
	"this is a weird\n\n\n\n\n\n\n        string\n",
}

func BenchmarkGetCounts(b *testing.B) {
	for i := range b.N {
		data := benchData[i%(len(benchData))]

		r := strings.NewReader(data)

		GetCounts(r)
	}
}

func BenchmarkGetCountsSinglePass(b *testing.B) {
	for i := range b.N {
		data := benchData[i%(len(benchData))]

		r := strings.NewReader(data)

		GetCountsSinglePass(r)
	}
}
