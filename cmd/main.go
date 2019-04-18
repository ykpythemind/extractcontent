package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	ext "github.com/ykpythemind/extractcontent"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "URL to extract")
	flag.Parse()

	if url == "" {
		// 標準入力がある
		os.Exit(extractStdin(os.Stdin, os.Stdout))
	} else {
		os.Exit(extractRemote(os.Stdout, url))
	}
}

func extractStdin(r io.Reader, w io.Writer) int {
	extractor := ext.NewExtractor(r, w, nil, false)
	if err := extractor.Extract(); err != nil {
		fmt.Fprintf(os.Stderr, "extract failed: %s\n", err)
		return 1
	}
	return 0
}

func extractRemote(w io.Writer, url string) int {
	if url == "" {
		fmt.Fprint(os.Stderr, "fail: URL is blank. \n")
		return 1
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err Get: %s\n", err)
		return 1
	}
	defer resp.Body.Close()

	extractor := ext.NewExtractor(resp.Body, w, nil, false)
	if err := extractor.Extract(); err != nil {
		fmt.Fprintf(os.Stderr, "extract failed: %s\n", err)
		return 1
	}
	return 0
}
