package extractcontent

import (
	"io"

	"golang.org/x/net/html"
)

// Sanitizer sanitize nodes and write results to io.Writer
type Sanitizer interface {
	Sanitize(*html.Node, io.Writer) error
}

type NoopSanitizer struct {
}

func (n *NoopSanitizer) Sanitize(node *html.Node, w io.Writer) error {
	return html.Render(w, node)
}
