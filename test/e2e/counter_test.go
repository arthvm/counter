package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func getCommand(args ...string) (*exec.Cmd, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, binName)

	cmd := exec.Command(path, args...)

	return cmd, nil
}

func TestStdin(t *testing.T) {
	cmd, err := getCommand()
	if err != nil {
		t.Fatal("couldn't get working directory:", err)
	}

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

func TestNoExist(t *testing.T) {
	cmd, err := getCommand("noexist.txt")
	if err != nil {
		t.Fatal("couldn't create command")
	}

	stderr := &bytes.Buffer{}
	stdout := &bytes.Buffer{}

	cmd.Stderr = stderr
	cmd.Stdout = stdout

	wantsStderr := "counter: open noexist.txt: no such file or directory\n"
	wantsStdout := ""

	err = cmd.Run()
	if err == nil {
		t.Log("command succeded when should have failed")
		t.Fail()
	}

	if err.Error() != "exit status 1" {
		t.Log("expected error of exit status 1. got:", err.Error())
		t.Fail()
	}

	if stderr.String() != wantsStderr {
		t.Log("stderr not match: wants:", wantsStderr, "got:", stderr)
		t.Fail()
	}

	if stdout.String() != wantsStdout {
		t.Log("stdout not match: wants:", wantsStdout, "got:", stdout)
		t.Fail()
	}
}
