package e2e

import (
	"fmt"
	"os"
	"testing"

	"github.com/arthvm/counter/test/assert"
)

func TestSingleFile(t *testing.T) {
	file, err := os.CreateTemp("", "counter-test-*")
	assert.NoError(t, err, "create file")
	defer os.Remove(file.Name())

	_, err = file.WriteString("foo bar baz\nbaz bar foo\none two three\n")
	assert.NoError(t, err, "write to file")

	err = file.Close()
	assert.NoError(t, err, "close file")

	cmd, err := getCommand(file.Name())
	assert.NoError(t, err, "get working dir")

	output, err := cmd.Output()
	assert.NoError(t, err, "run command")

	wants := fmt.Sprintf(" 3 9 38 %s\n", file.Name())
	assert.Equal(t, wants, string(output), "stdout is invalid")
}
