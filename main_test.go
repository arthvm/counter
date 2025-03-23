package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	wants := 5

	res := countWords([]byte(input))

	if res != wants {
		t.Fail()
	}
}
