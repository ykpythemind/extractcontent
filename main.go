package main

// https://blog.tottokug.com/entry/2017/12/09/233000
// https://github.com/tottokug/Trimmer/blob/master/com/tottokug/Trimmer.java

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// file, err := os.Open("testdata/test01.html")
	// if err != nil {
	// 	log.Fatalf("cant open file %s", err)
	// }
	// defer file.Close()

	// parse(file)

	if len(os.Args) != 2 {
		log.Fatal("args is wrong")
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("err Get: %s", err)
	}
	defer resp.Body.Close()

	parse(resp.Body)
}

func parse(r io.Reader) {
	nodes, err := html.Parse(r)
	if err != nil {
		log.Fatalf("err parsing: %s", err)
	}

	// TODO: cleaning node
	node := TrimNode(nodes, 0)

	debugNode(node)
	fmt.Printf("---------\n")

	fmt.Printf("もしかして本文は = %+v\n", getText(node))
}

// TrimNode は自分自身と子供達の中で一番強いノードを返す
func TrimNode(node *html.Node, score float64) *html.Node {
	maxScore := score
	strongNode := node

	children := getChildrenNode(node)
	for _, c := range children {
		if isSkippableNode(c) {
			continue
		}

		childScore := getScore(c)
		strongChild := TrimNode(c, childScore)
		strongScore := getScore(strongChild)
		if maxScore < strongScore {
			maxScore = strongScore
			strongNode = strongChild
		}
		fmt.Printf("TAG: %8s | SCORE: %6f\n", c.Data, childScore)
	}
	return strongNode
}

func getChildrenNode(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func getText(node *html.Node) string {
	var text []string
	if node.Type == html.TextNode && !isSkippableNode(node) {
		text = append(text, node.Data)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if isSkippableNode(c) {
			continue
		}
		if t := getText(c); t != "" {
			text = append(text, t)
		}
	}

	return strings.Join(text, "\n")
}

func getScore(node *html.Node) float64 {
	text := getTextLength(node)
	child := getChildrenCount(node)
	depth := getDepth(node)
	if child == 0 {
		child = 10000
	}

	fmt.Printf("TEXTLENGTH: %10d | CHILD: %6d | DEPTH: %6d\n", text, child, depth)

	tlen := float64(text)
	tmp := math.Sqrt((tlen * math.Sqrt(tlen)) * float64(depth) / math.Sqrt(float64(child*2)))
	score := math.Round(tmp)

	if node.DataAtom == atom.Section {
		score *= 1.5
	}

	return score
}

func getTextLength(node *html.Node) int {
	length := 0
	if isSkippableNode(node) {
		return 0
	}
	if node.Type == html.TextNode {
		length += len(node.Data) // TODO: 改行コードなどを置き換え
	} else {
		children := getChildrenNode(node)
		for _, c := range children {
			if c.DataAtom == atom.A {
				length += getTextLength(c) / 3 // a タグなどの文字数は軽くされるらしい
			} else {
				length += getTextLength(c)
			}
		}
	}
	return length
}

func getChildrenCount(node *html.Node) (count int) {
	children := getChildrenNode(node)
	for _, c := range children {
		if c.Type == html.ElementNode {
			if c.DataAtom == atom.Script || c.DataAtom == atom.Style || c.DataAtom == atom.Img || c.DataAtom == atom.Span {
				continue
			}
			count += getChildrenCount(c)
			count++
		}
	}

	return count
}

func getDepth(node *html.Node) (depth int) {
	p := node.Parent
	for ; p != nil; p = p.Parent {
		depth++
	}
	return
}

func isSkippableNode(node *html.Node) bool {
	if node.Type == html.TextNode {
		if node.Data == "" {
			log.Printf("Data is blank")
			return true
		}
	}
	if node.Type == html.CommentNode {
		return true
	}
	if node.DataAtom == atom.Script {
		return true
	}
	if node.DataAtom == atom.Style {
		return true
	}
	return false
}

func debugNode(node *html.Node) {
	fmt.Print("-------\n")
	fmt.Printf("Type = %+v\n", node.Type)
	fmt.Printf("Data = %+v\n", node.Data)
	fmt.Printf("DataAtom = %s\n", node.DataAtom.String())
	fmt.Printf("Namespace = %+v\n", node.Namespace)
	for _, a := range node.Attr {
		fmt.Printf("  attr = %+v\n", a)
	}
}
