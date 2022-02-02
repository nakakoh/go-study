package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

// 構造体埋め込みによる型の合成
func main() {
	{
		var cp ColoredPoint
		cp.X = 1
		fmt.Println(cp.Point.X) // 1
		cp.Point.Y = 2
		fmt.Println(cp.Y) // 2

		red := color.RGBA{255, 0, 0, 255}
		blue := color.RGBA{0, 0, 255, 255}
		var p = ColoredPoint{Point{1, 1}, red}
		var q = ColoredPoint{Point{5, 4}, blue}
		// coloredPoint型でもPoint型のメソッドを呼び出せる
		// PointのメソッドはColoredPointへ格上げ(promoted)されている
		// コンポジション(composition)
		// Pointをベースクラスとしてみなして、ColoredPointはサブクラスという考え方は誤り
		p.ScaleBy(2)
		q.ScaleBy(2)
		fmt.Println(p.Distance(q.Point)) // 10
		//p.Distance(q) // コンパイルエラー: Pointとしては q(ColoredPoint) を使えない
	}
	{
		red := color.RGBA{255, 0, 0, 255}
		blue := color.RGBA{0, 0, 255, 255}
		p := ColoredPoint2{&Point{1, 1}, red}
		q := ColoredPoint2{&Point{5, 4}, blue}
		fmt.Println(p.Distance(*q.Point)) // 5
		q.Point = p.Point                 // p と q はここで同じPointを共有する
		p.ScaleBy(2)
		fmt.Println(*p.Point, *q.Point) // {2 2} {2 2}
	}
	{
		var (
			mu      sync.Mutex // mapping を保護する
			mapping = make(map[string]string)
			Lookup  = func(key string) string {
				mu.Lock()
				v := mapping[key]
				mu.Unlock()
				return v
			}
		)
		mapping["hoge"] = "fuga"
		fmt.Println(Lookup("hoge"))
	}
	{
		var cache = struct {
			sync.Mutex
			mapping map[string]string
		}{
			mapping: make(map[string]string),
		}

		var Lookup = func(key string) string {
			cache.Lock()
			v := cache.mapping[key]
			cache.Unlock()
			return v
		}
		cache.mapping["hoge"] = "fuga"
		fmt.Println(Lookup("hoge"))
	}

}
