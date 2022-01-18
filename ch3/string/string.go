package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	{
		s := "hello, world"
		fmt.Println(len(s))     // 12
		fmt.Println(s[0], s[7]) // 104 119
		//fmt.Println(s[len(s)])  // panic!
		fmt.Println(s[0:5])            // hello
		fmt.Println(s[:5])             // hello
		fmt.Println(s[7:])             // world
		fmt.Println(s[:])              // hello, world
		fmt.Println("goodbye" + s[5:]) // hello, world
	}
	{
		s := "left foot"
		t := s
		s += ", right foot"
		fmt.Println(s) // left foot, right foot
		fmt.Println(t) // left foot
		//s[0] = 'L' // コンパイルエラー: s[0] には代入できない
	}

	// 生文字列リテラル (raw string literal)
	const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]`
	fmt.Println(GoUsage)

	// 全部 "世界"
	fmt.Println("世界")
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")

	fmt.Println(HasPrefix("世界", "世"))      // true
	fmt.Println(HasPrefix("世界", "界"))      // false
	fmt.Println(Contains("世界", "世"))       // true
	fmt.Println(Contains("世界", "界"))       // true
	fmt.Println(Contains("世界", "せ"))       // true
	fmt.Println(Contains("世界に一つだけ", "一つ")) // true

	{
		// len() はバイト数をカウントするので注意
		s := "Hello, 世界"
		fmt.Println(len(s))                    // 13
		fmt.Println(utf8.RuneCountInString(s)) // 9

		for i := 0; i < len(s); {
			r, size := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("%d\t%c\n", i, r)
			i += size
		}

		// 各文字のバイトインデックスとruneを表示
		for i, r := range "Hello, 世界" {
			fmt.Printf("%d\t%q\t%d\n", i, r, r)
		}

		// utf8.RuneCountInString(s) と同じ
		n := 0
		for _, _ = range s {
			n++
		}
		fmt.Println(n)

		// 上記と同じ
		n = 0
		for range s {
			n++
		}
		fmt.Println(n)
	}

	{
		// rune 変換
		s := "プログラム"
		// % x は16進数表記の2桁ごとの間に空白を挿入する
		fmt.Printf("% x\n", s) // e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0
		r := []rune(s)
		fmt.Printf("%x\n", r)        // [30d7 30ed 30b0 30e9 30e0]
		fmt.Println(string(r))       // プログラム
		fmt.Println(string(65))      // A, not 65
		fmt.Println(string(0x4eac))  // 京
		fmt.Println(string(1234567)) // �
	}
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
