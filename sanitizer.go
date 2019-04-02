package extractcontent

import (
	"bytes"
	"io"

	"github.com/microcosm-cc/bluemonday"
)

var strictPolicy = bluemonday.StrictPolicy()

// NOTE: id属性などは残ってしまう
var defaultPolicy = bluemonday.UGCPolicy()

// Sanitizer sanitize from reader and return results
type Sanitizer interface {
	Sanitize(io.Reader) *bytes.Buffer
}

type strictSanitizer struct {
}

type defaultSanitizer struct {
}

func (n *defaultSanitizer) Sanitize(r io.Reader) *bytes.Buffer {
	return defaultPolicy.SanitizeReader(r)
}

func (s *strictSanitizer) Sanitize(r io.Reader) *bytes.Buffer {
	return strictPolicy.SanitizeReader(r)
}
