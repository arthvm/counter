package e2e

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestStdin(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("couldn't get working directory:", err)
	}

	path := filepath.Join(dir, binName)

	cmd := exec.Command(path)

	output := &bytes.Buffer{}

	cmd.Stdin = strings.NewReader("one two three\n")
	cmd.Stdout = output

	if err := cmd.Run(); err != nil {
		t.Fatal("failed to run command")
	}

	wants := " 1 3 14\n"

	if wants != output.String() {
		t.Log("stdout is not correct: wanted:", wants, "got: ", output.String())
		t.Fail()
	}
}
