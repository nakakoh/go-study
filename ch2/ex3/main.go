package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"go-study.example.com/ch2/ex3/popcount"
)

func main() {
	var tmp string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("input: ")
	if scanner.Scan() {
		tmp = scanner.Text()
	}
	val, err := strconv.ParseUint(tmp, 0, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to ParseInt: %s", tmp)
	}
	fmt.Fprintf(os.Stdout, "%v\n", popcount.PopCount(val))
}
