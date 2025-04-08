package e2e

import (
	"fmt"
	"os"
	"testing"

	"github.com/arthvm/counter/test/assert"
)

func TestFlags(t *testing.T) {
	file, err := createFile("one two three four five\none two three\n")
	assert.NoError(t, err, "create file")
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
			assert.NoError(t, err, "get command")

			output, err := cmd.Output()
			assert.NoError(t, err, "run command")

			assert.Equal(t, tc.wants, string(output))
		})
	}
}
