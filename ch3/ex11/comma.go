package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1234"))
	fmt.Println(comma("5647532"))
	fmt.Println(comma("1234567890"))
	fmt.Println(comma("43489089270890432"))
	fmt.Println(comma("12.32"))
	fmt.Println(comma("124343.12323"))
}

func comma(s string) string {
	var decimalPart string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		decimalPart = s[dot:]
		s = s[:dot-1]
	}
	b := []byte(s)
	var buf bytes.Buffer
	for i, l := 0, len(b); i < l; i++ {
		if i > 0 && (l-i)%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(b[i])
	}
	buf.WriteString(decimalPart)
	return buf.String()
}
