package main

// run ../../ch1/fetch/fetch.go https://golang.org | go run count.go
import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "count: %v\n", err)
		os.Exit(1)
	}
	countMap := make(map[string]int)
	countTag(countMap, doc)

	for k, v := range countMap {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func countTag(countMap map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		countMap[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countTag(countMap, c)
	}
}
