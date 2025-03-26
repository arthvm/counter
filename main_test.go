package main_test

import (
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
			res := counter.CountWords(r)

			if res != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}
