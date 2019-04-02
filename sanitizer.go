package extractcontent

import (
	"bytes"
	"io"

	"github.com/microcosm-cc/bluemonday"
)

var policy = bluemonday.StrictPolicy()

// Sanitizer sanitize from reader and return results
type Sanitizer interface {
	Sanitize(io.Reader) (*bytes.Buffer, error)
}

// StrictSanitizer remove all tags
type StrictSanitizer struct {
}

type NoopSanitizer struct {
}

func (n *NoopSanitizer) Sanitize(r io.Reader) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(r)
	return buf, err
}

func (s *StrictSanitizer) Sanitize(r io.Reader) (*bytes.Buffer, error) {
	b := policy.SanitizeReader(r)
	return b, nil
}
