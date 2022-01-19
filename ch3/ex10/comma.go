package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234"))
	fmt.Println(comma("5647532"))
	fmt.Println(comma("1234567890"))
	fmt.Println(comma("43489089270890432"))
}

func comma(s string) string {
	b := []byte(s)
	var buf bytes.Buffer
	for i, l := 0, len(b); i < l; i++ {
		if i > 0 && (l-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(b[i])
	}
	return buf.String()
}
