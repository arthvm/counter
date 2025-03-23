package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	wants := 5

	res := CountWords([]byte(input))

	if res != wants {
		t.Logf("expected: %d got: %d", wants, res)
		t.Fail()
	}

	input = ""
	wants = 0

	res = CountWords([]byte(input))

	if res != wants {
		t.Logf("expected: %d got: %d", wants, res)
		t.Fail()
	}

	input = " "
	wants = 0

	res = CountWords([]byte(input))

	if res != wants {
		t.Logf("expected: %d got: %d", wants, res)
		t.Fail()
	}
}
