package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// go run main.go http://gopl.io
// go run main.go https://golang.org/doc/effective_go.html
// go run main.go https://golang.org/doc/gopher/frontpage.png

func main() {
	url := os.Args[1]
	err := title(url)
	if err != nil {
		log.Print(err)
	}
}

// soleTitle は doc 中の最初の空ではない title 要素のテキストと、
// title 要素が一つだけでなかったらエラーを返します。
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// パニックなし
		case bailout{}:
			// 「予期された」パニック
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // 予期しないパニック; パニックを維持する
		}
	}()

	// 二つ以上の空ではない title を見つけたら再帰から抜け出させる。
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // 複数の title 要素
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// defer は遅延関数呼び出し。関数が完了するまで呼び出しを遅延(deferred)される
	// returnを実行したり、関数の最後に到達したり、パニックによる異常完了でも実行してくれる
	defer resp.Body.Close()

	// Content-Type が HTML (例: "text/html; charset=utf-8") であるかを検査する。
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
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
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
