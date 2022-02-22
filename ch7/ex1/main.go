package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	c := 0
	for scanner.Scan() {
		c++
	}
	*w += WordCounter(c)
	return c, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	c := 0
	for scanner.Scan() {
		c++
	}
	*l += LineCounter(c)
	return c, nil
}

func main() {
	var w WordCounter
	var l LineCounter
	input := `kore ha test desu.
	are mo test desu.
	yorosiku.`

	w.Write([]byte(input))
	fmt.Println(w)
	l.Write([]byte(input))
	fmt.Println(l)
}
