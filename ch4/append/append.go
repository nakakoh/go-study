package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
	// 以下も同じ
	fmt.Printf("%q\n", []rune("Hello, 世界"))

	// appendInt
	{
		var x, y []int
		for i := 0; i < 10; i++ {
			y = appendInt(x, i)
			fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
			x = y
		}
	}
	// appendInt2
	{
		var x, y []int
		for i := 0; i < 10; i++ {
			y = appendInt2(x, 1, 2)
			fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
			x = y
		}
	}

	// push, pop
	{
		stack := []string{}
		stack = append(stack, "hoge") // push
		fmt.Println(stack)
		top := stack[len(stack)-1]
		fmt.Println(top)
		stack = stack[:len(stack)-1] // pop
		fmt.Println(stack)
	}
	// remove
	{
		s := []int{5, 6, 7, 8, 9}
		fmt.Println(remove(s, 2)) // [5,6,8,9]
		fmt.Println(s)
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 拡大する余地がある。スライスを拡張する。
		z = x[:zlen]
	} else {
		// 十分な領域がない。新たな配列を割り当てる。
		// 計算量を線形に均すために倍に拡大する。
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 組み込み関数
	}
	z[len(x)] = y
	return z
}

// 複数append対応
func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// 拡大する余地がある。スライスを拡張する。
		z = x[:zlen]
	} else {
		// 十分な領域がない。新たな配列を割り当てる。
		// 計算量を線形に均すために倍に拡大する。
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 組み込み関数
	}
	copy(z[len(x):], y)
	return z
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
