package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/arthvm/counter/test/assert"
)

func TestMultipleFiles(t *testing.T) {
	fileA, err := createFile("one two three four five\n")
	assert.NoError(t, err, "create fileA")

	defer os.Remove(fileA.Name())

	fileB, err := createFile("foo bar baz\n\n")
	assert.NoError(t, err, "create fileB")

	defer os.Remove(fileB.Name())

	fileC, err := createFile("")
	assert.NoError(t, err, "create fileC")

	defer os.Remove(fileC.Name())

	cmd, err := getCommand(fileA.Name(), fileB.Name(), fileC.Name())
	assert.NoError(t, err, "create command")

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	err = cmd.Run()
	assert.NoError(t, err, "run command")

	wants := fmt.Sprintf(` 1 5 24 %s
 2 3 13 %s
 0 0  0 %s
 3 8 37 totals
`, fileA.Name(), fileB.Name(), fileC.Name())

	res := stdout.String()
	assert.Equal(t, wants, res)
}
