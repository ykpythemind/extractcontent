package extractcontent

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func newExtractor(r io.Reader) (*Extractor, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	return NewExtractor(r, buf, nil, testing.Verbose()), buf
}

func TestExtractFile01(t *testing.T) {
	file, err := os.Open("testdata/test01.html")
	if err != nil {
		t.Fatalf("cant open file %s", err)
	}
	defer file.Close()

	e, buf := newExtractor(file)

	err = e.Extract()
	if err != nil {
		t.Fatalf("failed to extract: %s", err)
	}

	t.Logf("test01 result: %s", buf)

	str := buf.String()
	prefix := "<section"
	if !strings.HasPrefix(str, prefix) {
		t.Fatalf("results must start with: %s", prefix)
	}
}
