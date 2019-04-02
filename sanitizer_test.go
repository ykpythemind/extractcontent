package extractcontent

import (
	"testing"
)

func TestPolicy(t *testing.T) {
	html := `<html><body><article><img src="https://google.com"><iframe></iframe><p id="hoge">fuga</p></article></body></html>`

	result := defaultPolicy.Sanitize(html)
	if result != `<article><img src="https://google.com"><p id="hoge">fuga</p></article>` {
		t.Fatalf("sanitize result not match: result = %+v\n", result)
	}
}
