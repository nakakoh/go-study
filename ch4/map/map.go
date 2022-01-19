package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		ages := make(map[string]int) // 文字列から int へのマッピング
		fmt.Println(ages)
	}

	// マップリテラル (map literal) を使ってMapの作成
	{
		ages := map[string]int{
			"alice":   31,
			"charlie": 34,
		}
		fmt.Println(ages)
	}
	// 上記は以下と同じ
	{
		ages := make(map[string]int)
		ages["alice"] = 31
		ages["charlie"] = 34
		fmt.Println(ages)
	}
	{
		ages := map[string]int{
			"alice":   31,
			"charlie": 34,
		}
		// キーを指定して取り出す
		fmt.Println(ages["alice"]) // 31
		// 要素の削除
		delete(ages, "alice")
		fmt.Println(ages) // charlieだけになる
		// 存在しないキーでも動作する（ゼロ値扱い）
		fmt.Println(ages["bob"])
		ages["bob"] = ages["bob"] + 1 // ages["bob"] += 1 ; ages["bob"]++ でもOK
		fmt.Println(ages["bob"])
		//_ = &ages["bob"] // コンパイルエラー: マップの要素のアドレスは得られない

		for name, age := range ages {
			fmt.Printf("%s\t%d\n", name, age)
		}
	}
	// ソート
	{
		fmt.Println("--- sort ---")
		ages := map[string]int{"alice": 31, "charlie": 34, "tom": 20, "bob": 1}
		var names []string
		for name := range ages {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("%s\t%d\n", name, ages[name])
		}
	}
	// マップ型のゼロ値はnil
	{
		var ages map[string]int
		fmt.Println(ages == nil)    // true
		fmt.Println(len(ages) == 0) // true
		//ages["carol"] = 21          // パニック: nilマップのエントリへの代入
		age, ok := ages["bob"]
		fmt.Println(age, ok)
	}
	// マップもスライス同様互いに比較できない。nilだけは比較できる
	{
		fmt.Println("--- compare ---")
		equal(map[string]int{"A": 0}, map[string]int{"B": 42}) // false
		ages := map[string]int{"alice": 31, "charlie": 34, "tom": 20, "bob": 1}
		ages2 := map[string]int{"alice": 31, "charlie": 34, "tom": 20, "bob": 1}
		ages3 := map[string]int{"alice": 31, "tom": 20, "bob": 1}
		fmt.Println(equal(ages, ages2)) // true
		fmt.Println(equal(ages, ages3)) // false
	}

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
