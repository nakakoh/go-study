package main

import "fmt"

func main() {
	{
		var a [3]int             // 3この整数の配列
		fmt.Println(a[0])        // 最初の要素を表示する
		fmt.Println(a[len(a)-1]) // 最後の要素である a[2] を表示する

		// インデックスと要素を表示
		for i, v := range a {
			fmt.Printf("%d %d\n", i, v)
		}
		// 要素だけを表示
		for _, v := range a {
			fmt.Printf("%d\n", v)
		}
	}

	// 配列リテラル
	{
		// 初期化時に値を宣言
		var q [3]int = [3]int{1, 2, 3}
		var r [3]int = [3]int{1, 2}
		for i, v := range q {
			fmt.Printf("%d %d\n", i, v)
		}
		for i, v := range r {
			fmt.Printf("%d %d\n", i, v)
		}
		// ... を使えば初期化子の数でサイズを決定
		{
			q := [...]int{1, 2, 3}
			fmt.Printf("%T\n", q)
		}
		// 一度サイズが確定したものに別なサイズのものを代入できない（異なる型のため）
		{
			//q := [3]int{1, 2, 3}
			//q = [4]int{1, 2, 3, 4} // コンパイルエラー: [4]int を [3]int に代入できない
		}
	}
	// index と値の組のリスト指定
	{
		type Currency int
		const (
			USD Currency = iota
			EUR
			GBP
			RMB
		)
		symbol := [...]string{USD: "$", EUR: "e", GBP: "£", RMB: "¥"}
		fmt.Println(RMB, symbol[RMB]) // 3 ¥
	}
	// 100個の要素を持ち、最後の要素だけが -1
	{
		r := [...]int{99: -1}
		for i, v := range r {
			fmt.Printf("[%d]%d", i, v)
		}
		fmt.Println()
	}
	// 配列の比較はすべての対応する要素が等しいかを判定
	// 型が違うと比較できない
	{
		a := [2]int{1, 2}
		b := [...]int{1, 2}
		c := [2]int{1, 3}
		fmt.Println(a == b, a == c, b == c) // true false false
		//d := [3]int{1, 2}
		//fmt.Println(a == d) // コンパイルエラー: 比較できない [2]int == [3]int
	}

	{
		// goでは関数のパラメータにはコピーが渡されるので参照を渡すときはポインタで
		a := [32]byte{1, 1}
		fmt.Printf("%v\n", a)
		zero(&a)
		fmt.Printf("%v\n", a)
	}

}

// 配列を全部0にする
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
		// 以下も同じ
		//*ptr = [32]byte{}
	}
}
