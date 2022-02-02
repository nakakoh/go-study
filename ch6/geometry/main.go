package main

import (
	"fmt"

	"go-study.example.com/ch6/geometry/geometry"
)

func main() {
	{
		p := geometry.Point{1, 2}
		q := geometry.Point{4, 6}
		fmt.Println(geometry.Distance(p, q)) // 5, 関数呼び出し
		fmt.Println(p.Distance(q))           // 5, メソッド呼び出し
	}

	{
		perim := geometry.Path{
			{1, 1},
			{5, 1},
			{5, 4},
			{1, 1},
		}
		fmt.Println(perim.Distance()) // 12
	}

	// ポインタレシーバーを持つメソッドの呼び方
	{
		// 1. アドレスで初期化して呼び出し
		r := &geometry.Point{1, 2}
		r.ScaleBy(2)
		fmt.Println(*r) // {2, 4}

		// 2. アドレスを別途変数に代入して呼び出し
		p := geometry.Point{1, 2}
		pptr := &p
		pptr.ScaleBy(2)
		fmt.Println(p) // {2, 4}

		// 3. 呼び出し時にアドレスを指定
		q := geometry.Point{1, 2}
		(&q).ScaleBy(2)
		fmt.Println(q) // {2, 4}

		// 4. アドレス指定しなくてもよい(つまりわざわざ2,3の書き方をしなくてもよい)
		// コンパイラが変数に対して暗黙的に &s してくれる
		s := geometry.Point{1, 2}
		s.ScaleBy(2)
		fmt.Println(s) // {2, 4}

		// ただし以下はコンパイルエラー: Pointリテラルのアドレスは得られない
		//geometry.Point{1 ,2}.ScaleBy(2)

		// Point.DistanceのようなPointのメソッドを *Pointレシーバーで呼び出すことはできる
		// アドレスから値を得る方法があるから
		fmt.Println(pptr.Distance(q))
		fmt.Println((*pptr).Distance(q))

		// (1) レシーバ引数がレシーバーパラメータと同じ型である。
		// 例えば、両方が型T、もしくは両方が型*T のとき
		geometry.Point{1, 2}.Distance(q) // Point
		pptr.ScaleBy(2)                  // *Point

		// (2) レシーバ引数が型T の変数で、レシーバパラメータが型*T である。
		// コンパイラは暗黙的にその変数のアドレスを得る
		p.ScaleBy(2) // 暗黙的に(&p)

		// (3) レシーバ引数の型が *T であり、レシーバパラメータが型Tである
		// コンパイラは暗黙的にレシーバの指す値を参照する（つまりその値をロードする）
		pptr.Distance(q) // 暗黙的に(*pptr)
	}
}

// IntList は整数のリンクリスト
// nil の*IntList は空リストを表します
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum はリスト要素の合計値を返します
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
