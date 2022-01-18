package main

import (
	"fmt"
)

func main() {
	fmt.Println(-5 % 3)
	fmt.Println(-5 % -3)
	fmt.Println(5 / 4)
	fmt.Println(7 / 4)

	{
		var u uint8 = 255
		fmt.Println(u, u+1, u*u) // "255 0 1"
		var i int8 = 127
		fmt.Println(i, i+1, i*i) // "127 -128 1"
		fmt.Println("--------------")
	}

	{
		var x uint8 = 1<<1 | 1<<5
		var y uint8 = 1<<1 | 1<<2
		fmt.Printf("%08b\n", x)
		fmt.Printf("%08b\n", y)

		fmt.Printf("%08b\n", x&y)
		fmt.Printf("%08b\n", x|y)
		fmt.Printf("%08b\n", x^y)
		fmt.Printf("%08b\n", x&^y)

		for i := uint(0); i < 8; i++ {
			if x&(1<<i) != 0 {
				fmt.Println(i)
			}
		}
		fmt.Printf("%08b\n", x<<1)
		fmt.Printf("%08b\n", x>>1)
	}

	// len() は int型を返す(uintだとやばい)
	fmt.Println("----- medals -----")
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	fmt.Println("----- convert -----")
	var apples int32 = 1
	var oranges int16 = 2
	// 以下はコンパイルエラー
	//var compote int = apples + oranges // コンパイルエラー
	// 型を変換する
	var compote = int(apples) + int(oranges)
	fmt.Println(compote)

	// 型変換は精度を失うかもしれない
	{
		f := 3.141 // float64
		i := int(f)
		fmt.Println(f, i) // 3.141 3
		f = 1.99
		fmt.Println(int(f)) // 1
	}

	{
		o := 0666
		fmt.Printf("%d %[1]o %#[1]o\n", o)
		x := int64(0xdeadbeef)
		fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	}

	// rune ルーン
	{
		ascii := 'a'
		unicode := '国'
		newline := '\n'
		fmt.Printf("%d %[1]c %[1]q\n", ascii)   // 97 a 'a'
		fmt.Printf("%d %[1]c %[1]q\n", unicode) // 22269 国 '国'
		fmt.Printf("%d %[1]q\n", newline)       // 10 '\n'
	}

}
