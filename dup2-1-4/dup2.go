package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileSets := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileSets)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileSets)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				filename := fileSets[line]
				fmt.Printf("%d\t%s\t%s\n", n, line, filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileSets map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileSets[input.Text()] += f.Name()
	}
	// 注意: input.Err() からのエラーの可能性を無視している
}
