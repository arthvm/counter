package e2e

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
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

	wants := map[string]string{
		fileA.Name(): fmt.Sprintf(" 1 5 24 %s", fileA.Name()),
		fileB.Name(): fmt.Sprintf(" 2 3 13 %s", fileB.Name()),
		fileC.Name(): fmt.Sprintf(" 0 0  0 %s", fileC.Name()),
		"totals":     " 3 8 37 totals",
	}

	scanner := bufio.NewScanner(stdout)

	checkedWants := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 0 {
			t.Log("encountered empty line")
			t.Fail()
			continue
		}

		filename := fields[len(fields)-1]

		lineWants, ok := wants[filename]
		if !ok {
			t.Logf("no wants for filename: %s", filename)
			t.Fail()
		}

		checkedWants++

		if line != lineWants {
			t.Logf("line does not match: got: %s, wants: %s", line, lineWants)
			t.Fail()
		}
	}

	if checkedWants != len(wants) {
		t.Logf("only checked: %d expected to check: %d", checkedWants, len(wants))
		t.Fail()
	}
}
