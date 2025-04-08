package assert

import (
	"reflect"
	"strings"
	"testing"
)

func Equal(t *testing.T, wants any, got any, msg ...string) {
	t.Helper()

	if reflect.DeepEqual(wants, got) {
		return
	}

	msgStrr := strings.Join(msg, ":")

	t.Logf("%s\nexpected: %v \n     got: %v\n", msgStrr, wants, got)
	t.Fail()
}

func NoError(t *testing.T, err error, msg ...string) {
	t.Helper()

	if err == nil {
		return
	}

	msgStrr := strings.Join(msg, ":")
	t.Fatalf("failed to %s\n %v", msgStrr, err)
}
