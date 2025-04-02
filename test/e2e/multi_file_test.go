package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestMultipleFiles(t *testing.T) {
	fileA, err := createFile("one two three four five\n")
	if err != nil {
		t.Fatal("could not create fileA:", err)
	}

	defer os.Remove(fileA.Name())

	fileB, err := createFile("foo bar baz\n\n")
	if err != nil {
		t.Fatal("could not create fileB:", err)
	}

	defer os.Remove(fileB.Name())

	fileC, err := createFile("")
	if err != nil {
		t.Fatal("could not create fileC:", err)
	}

	defer os.Remove(fileC.Name())

	cmd, err := getCommand(fileA.Name(), fileB.Name(), fileC.Name())
	if err != nil {
		t.Fatal("could not create command:", err)
	}

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		t.Fatal("failed to run command:", err)
	}

	wants := fmt.Sprintf(` 1 5 24 %s
 2 3 13 %s
 0 0  0 %s
 3 8 37 totals
`, fileA.Name(), fileB.Name(), fileC.Name())

	res := stdout.String()
	if wants != res {
		t.Logf("expected: %s got: %s", wants, res)
		t.Fail()
	}
}
