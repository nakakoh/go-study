package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// go run outline.go https://golang.org
func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "getting %s: %s", url, resp.Status)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %s as HTML: %v", url, err)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		var sb strings.Builder
		for _, attr := range n.Attr {
			fmt.Fprintf(&sb, " %s=%s", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s />\n", depth*2, "", n.Data, sb.String())
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, sb.String())
			depth++
		}
	case html.TextNode:
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
