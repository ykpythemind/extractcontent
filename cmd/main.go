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

	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		panic(fmt.Sprintf("file.Stat() %s", err))
	}

	size := fi.Size()
	if size > 0 {
		// 標準入力がある
		os.Exit(ExtractStdin(os.Stdin, os.Stdout))
	} else {
		os.Exit(ExtractRemote(os.Stdout, url))
	}
}

func ExtractStdin(r io.Reader, w io.Writer) int {
	extractor := ext.NewExtractor(r, w, nil, false)
	if err := extractor.Extract(); err != nil {
		fmt.Fprintf(os.Stderr, "extract failed: %s\n", err)
		return 1
	}
	return 0
}

func ExtractRemote(w io.Writer, url string) int {
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
