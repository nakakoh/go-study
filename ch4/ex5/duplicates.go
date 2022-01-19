package main

import "fmt"

func removeAdjacentDuplicates(strings []string) []string {
	// 重複ない
	if len(strings) <= 1 {
		return strings
	}

	prev := strings[0]
	i := 1
	for _, s := range strings {
		if s != prev {
			prev = s
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	a := []string{"foo", "buz", "bar"}
	fmt.Println(removeAdjacentDuplicates(a))
	b := []string{"foo", "buz", "buz", "bar", "bar", "foo", "bar", "foo", "foo"}
	fmt.Println(removeAdjacentDuplicates(b))
}
