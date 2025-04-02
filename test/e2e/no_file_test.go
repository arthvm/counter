package e2e

import (
	"bytes"
	"testing"
)

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
