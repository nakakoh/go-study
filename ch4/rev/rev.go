package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	// 2つだけ左へsを回転させる
	reverse(s[:2]) // [1 0 2 3 4 5]
	fmt.Println(s)
	reverse(s[2:]) // [1 0 5 4 3 2]
	fmt.Println(s)
	reverse(s) // [2 3 4 5 0 1]
	fmt.Println(s)
}
