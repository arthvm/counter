package e2e

import (
	"bytes"
	"testing"

	"github.com/arthvm/counter/test/assert"
)

func TestNoExist(t *testing.T) {
	cmd, err := getCommand("noexist.txt")
	assert.NoError(t, err, "create command")

	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	wantsStderr := "counter: open noexist.txt: no such file or directory\n"
	wantsStdout := ""

	output, err := cmd.Output()
	if err == nil {
		t.Log("command succeded when should have failed")
		t.Fail()
	}

	if err.Error() != "exit status 1" {
		t.Log("expected error of exit status 1. got:", err.Error())
		t.Fail()
	}

	assert.Equal(t, wantsStderr, stderr.String(), "stderr doesn't match")
	assert.Equal(t, wantsStdout, string(output), "stdout doesn't match")
}
