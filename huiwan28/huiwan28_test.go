package huiwan28

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	path := "/Users/miaojingyi/Documents/dev/go/src/tty28/index.html"

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("os.Open() Failure :: %s", err.Error())
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		t.Fatalf("io.Copy() Failure :: %s", err.Error())
	}

	t.Log(buf.String())
}
