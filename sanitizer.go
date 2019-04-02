package extractcontent

import (
	"bytes"
	"io"

	"github.com/microcosm-cc/bluemonday"
)

var strictPolicy = bluemonday.StrictPolicy()
var policy = bluemonday.UGCPolicy()

// Sanitizer sanitize from reader and return results
type Sanitizer interface {
	Sanitize(io.Reader) *bytes.Buffer
}

// StrictSanitizer remove all tags
type StrictSanitizer struct {
}

type DefaultSanitizer struct {
}

func (n *DefaultSanitizer) Sanitize(r io.Reader) *bytes.Buffer {
	return policy.SanitizeReader(r)
}

func (s *StrictSanitizer) Sanitize(r io.Reader) *bytes.Buffer {
	return strictPolicy.SanitizeReader(r)
}
