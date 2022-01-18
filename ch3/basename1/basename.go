package main

import "fmt"

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
}

func basename(s string) string {
	// 最後の'/' とsの前の全てを破棄する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 最後の'.' より前のすべてを保持する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
