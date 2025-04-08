package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestFlags(t *testing.T) {
	file, err := createFile("one two three four five\none two three\n")
	if err != nil {
		t.Fatal("could not create file")
	}
	defer os.Remove(file.Name())

	testCases := []struct {
		name  string
		wants string
		flags []string
	}{
		{
			name:  "line flag",
			flags: []string{"-l"},
			wants: fmt.Sprintf(" 2 %s\n", file.Name()),
		},
		{
			name:  "bytes flag",
			flags: []string{"-c"},
			wants: fmt.Sprintf(" 38 %s\n", file.Name()),
		},
		{
			name:  "words flag",
			flags: []string{"-w"},
			wants: fmt.Sprintf(" 8 %s\n", file.Name()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputs := append(tc.flags, file.Name())

			cmd, err := getCommand(inputs...)
			if err != nil {
				t.Log("failed to get command:", err)
				t.Fail()
			}

			stdout := &bytes.Buffer{}
			cmd.Stdout = stdout

			if err := cmd.Run(); err != nil {
				t.Log("failed to run command:", err)
				t.Fail()
			}

			output := stdout.String()

			if output != tc.wants {
				t.Logf("output did not match got: %s wants: %s", output, tc.wants)
				t.Fail()
			}
		})
	}
}
