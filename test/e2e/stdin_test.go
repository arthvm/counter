package e2e

import (
	"strings"
	"testing"

	"github.com/arthvm/counter/test/assert"
)

func TestStdin(t *testing.T) {
	cmd, err := getCommand()
	assert.NoError(t, err, "get working dir")

	cmd.Stdin = strings.NewReader("one two three\n")

	output, err := cmd.Output()
	assert.NoError(t, err, "run command")

	wants := " 1 3 14\n"
	assert.Equal(t, wants, string(output), "stdout is invalid")
}
