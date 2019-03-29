package extractcontent

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func newExtractContent(r io.Reader) (*ExtractContent, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	return NewExtractContent(r, buf, true), buf
}

func TestExtractFile01(t *testing.T) {
	file, err := os.Open("testdata/test01.html")
	if err != nil {
		t.Fatalf("cant open file %s", err)
	}
	defer file.Close()

	e, buf := newExtractContent(file)

	err = e.Extract()
	if err != nil {
		t.Fatalf("failed to extract: %s", err)
	}

	t.Logf("test01 result: %s", buf)

	str := buf.String()
	prefix := "アポ電"
	if !strings.HasPrefix(str, prefix) {
		t.Fatalf("results must start with: %s", prefix)
	}
}
