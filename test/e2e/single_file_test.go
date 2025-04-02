package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestSingleFile(t *testing.T) {
	file, err := os.CreateTemp("", "counter-test-*")
	if err != nil {
		t.Fatal("couldn't create temp file")
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString("foo bar baz\nbaz bar foo\none two three\n")
	if err != nil {
		t.Fatal("couldn't write to temp file")
	}

	err = file.Close()
	if err != nil {
		t.Fatal("couldn't close file")
	}

	cmd, err := getCommand(file.Name())
	if err != nil {
		t.Fatal("couldn't get working directory:", err)
	}

	output := &bytes.Buffer{}
	cmd.Stdout = output

	if err := cmd.Run(); err != nil {
		t.Fatal("failed to run command:", err)
	}

	wants := fmt.Sprintf(" 3 9 38 %s\n", file.Name())

	if output.String() != wants {
		t.Log("stdout not expected. wants: ", wants, "got:", output)
		t.Fail()
	}
}
