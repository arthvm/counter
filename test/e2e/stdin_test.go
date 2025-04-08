package e2e

import (
	"bytes"
	"strings"
	"testing"

	"github.com/arthvm/counter/test/assert"
)

func TestStdin(t *testing.T) {
	cmd, err := getCommand()
	assert.NoError(t, err, "get working dir")

	output := &bytes.Buffer{}

	cmd.Stdin = strings.NewReader("one two three\n")
	cmd.Stdout = output

	err = cmd.Run()
	assert.NoError(t, err, "run command")

	wants := " 1 3 14\n"
	assert.Equal(t, wants, output.String(), "stdout is invalid")
}
