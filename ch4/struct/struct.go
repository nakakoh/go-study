package main

import (
	"fmt"
	"image/gif"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// 型が同じならまとめることもできる
type Employee2 struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

// フィールドの順序は型の同一性にとって重要

func main() {
	var dilbert Employee

	dilbert.Salary -= 5000 // ほとんどコードを書かないので降格ｗ
	// フィールドのアドレスからポインタアクセス
	position := &dilbert.Position
	*position = "senior " + *position // Elbonia へアウトソースして昇進

	// ポインタでもドット表示可能
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	// 以下も同じ
	//(*employeeOfTheMonth).Position += " (proactie team player)"

	// 空構造体でboolの代わりに使う人もいるが、メモリの節約はわずかなので使うメリットはあまりない
	seen := make(map[string]struct{}) // 文字列のセット
	s := "first"
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}

	// 構造体リテラル(struct literal)
	{
		type Point struct{ X, Y int }
		// 後でフィールドを追加されたり順序を変えられたりすることに対して脆弱になる
		// 全てのフィールドに対して正しい順序で値を指定する必要がある
		// image.Point {x, y} color.RGBA{red, green, blue, alpha}などフィールド順が明らかな場合のみ
		// 使われる傾向にある
		_ = Point{1, 2}

		// 以下のようにフィールド名と値を指定して使うことのほうが多い
		// フィールドが省略された場合はゼロ値
		_ = gif.GIF{LoopCount: 10}

		// 他のパッケージから公開されてない識別子（先頭が大文字でないフィールド）を初期化時にも参照できない
		/*
			package p
			type T struct { a, b int } // a と b は公開されていない
			package q
			import "p"
			var _ = p.T{a: 1, b: 2} // コンパイルエラー: a と b を参照できない
			var _ = p.T{1, 2}       // コンパイルエラー: a と b を参照できない
		*/

		// 大きな構造体は、効率性のためにたいていポインタを使って間接的に関数に渡されたり、
		// 関数から返されたりする
		mike := Employee{ID: 1, Name: "mike", Salary: 10000}
		fmt.Println(mike)
		fmt.Println(Bonus(&mike, 300))
		AwardAnnualRaise(&mike) // 昇給！
		fmt.Println(mike)

		pp := &Point{1, 2}
		fmt.Println(pp, *pp)
	}

	// 構造体の比較
	{
		type Point struct{ X, Y int }
		p := Point{1, 2}
		q := Point{2, 1}
		fmt.Println(p.X == q.X && p.Y == q.Y) // false
		fmt.Println(p == q)                   // false

		// 比較可能な構造体型はマップのキー型として使える
		type address struct {
			hostname string
			port     int
		}
		hits := make(map[address]int)
		hits[address{"golang.orc", 443}]++
	}

	{
		type Circle struct {
			X, Y, Radius int
		}
		type Wheel struct {
			X, Y, Radius, Spokes int
		}
		var w Wheel
		w.X = 8
		w.Y = 8
		w.Radius = 5
		w.Spokes = 20
	}

	// 共通部分をくくりだすほうが便利になるがアクセスが面倒
	{
		type Point struct {
			X, Y int
		}
		type Circle struct {
			Center Point
			Radius int
		}
		type Wheel struct {
			Circle Circle
			Spokes int
		}
		var w Wheel
		w.Circle.Center.X = 8
		w.Circle.Center.Y = 8
		w.Circle.Radius = 5
		w.Spokes = 20
	}
	// 構造体埋め込み (struct embedding) による 無名フィールド (anonymous field)
	{
		type Point struct {
			X, Y int
		}
		type Circle struct {
			Point  // 無名フィールド
			Radius int
		}
		type Wheel struct {
			Circle // 無名フィールド
			Spokes int
		}
		// 無名フィールドの埋め込みによって、中間の名前を書くことなく参照できる
		// 無名フィールドとは言ってるが、名前を持っているので中間の名前を書いてアクセスも可能
		var w Wheel
		w.X = 8      // w.Circle.Point.X = 8 と同じ
		w.Y = 8      // w.Circle.Point.Y = 8 と同じ
		w.Radius = 5 // w.Circle.Radius = 5 と同じ
		w.Spokes = 20
		// ただし、構造体リテラルではちゃんと構造体も宣言しないとだめ
		//w = Wheel{8, 8, 5, 20} // コンパイルエラー: 不明フィールド
		//w = Wheel{X: 8,Y: 8, Radius: 5, Spokes: 20} // コンパイルエラー: 不明フィールド
		w = Wheel{Circle{Point{X: 8, Y: 8}, 5}, 20}
	}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
