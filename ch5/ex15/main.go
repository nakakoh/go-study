package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(max(1, 4, 2, 3))
	fmt.Println(min(1, 4, 2, 3))
	fmt.Println(max())
	fmt.Println(min())
	fmt.Println(max2(1, 4, 2, 3))
	fmt.Println(min2(1, 4, 2, 3))
	fmt.Println(max2(2))
	fmt.Println(min2(3))
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no args found")
	}
	sort.Sort(sort.IntSlice(vals))
	return vals[len(vals)-1], nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no args found")
	}
	sort.Sort(sort.IntSlice(vals))
	return vals[0], nil
}

func max2(v int, vals ...int) int {
	vals = append(vals, v)
	sort.Sort(sort.IntSlice(vals))
	return vals[len(vals)-1]
}

func min2(v int, vals ...int) int {
	vals = append(vals, v)
	sort.Sort(sort.IntSlice(vals))
	return vals[0]
}
