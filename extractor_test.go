package extractcontent

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"
)

func newExtractor(r io.Reader) (*Extractor, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	return NewExtractor(r, buf, nil, testing.Verbose()), buf
}

func prepareFile(name string) (*os.File, func()) {
	file, err := os.Open(path.Join("testdata", name))
	if err != nil {
		panic(fmt.Sprintf("cant open file %s", err))
	}

	return file, func() {
		file.Close()
	}
}

func TestExtractFile01(t *testing.T) {
	file, fn := prepareFile("test01.html")
	defer fn()

	extractor, buf := newExtractor(file)

	err := extractor.Extract()
	if err != nil {
		t.Fatalf("failed to extract: %s", err)
	}

	str := buf.String()
	t.Logf("test01 result: %s", str)

	prefix := "<section"
	if !strings.HasPrefix(str, prefix) {
		t.Fatalf("results must start with: %s", prefix)
	}
}

func TestExtractFile02(t *testing.T) {
	file, fn := prepareFile("test02_wikipedia.html")
	defer fn()

	extractor, buf := newExtractor(file)

	err := extractor.Extract()
	if err != nil {
		t.Fatalf("failed to extract: %s", err)
	}

	str := buf.String()
	t.Logf("test01 result: %s", str)

	prefix := "<section"
	if !strings.HasPrefix(str, prefix) {
		t.Fatalf("results must start with: %s", prefix)
	}
}
