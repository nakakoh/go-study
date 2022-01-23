package main

import "fmt"

func main() {
	type Point struct{ X, Y int }

	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	// #アドヴァーブでGo構文と似た形式で表示してくれる
	fmt.Printf("%#v\n", w)

	// 上記と同じ
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // 注意: 終わりのカンマは必要（Radiusも）
	}

	w.X = 42
	fmt.Printf("%#v\n", w)
}
