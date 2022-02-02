package main

import (
	"fmt"
	"math"
	"time"
)

type P *int

//func (P) f() { /* ... */ } // コンパイルエラー: 不正なレシーバー型

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {

	// メソッド値(method value)
	{
		p := Point{1, 2}
		q := Point{4, 6}

		distanceFromP := p.Distance        // メソッド値
		fmt.Println(distanceFromP(q))      // 5
		var origin Point                   // {0, 0}
		fmt.Println(distanceFromP(origin)) // 2.23606797749979 (√5)

		scaleP := p.ScaleBy // メソッド値
		scaleP(2)           // p は (2,4) になる
		fmt.Println(p)
		scaleP(3) // p は (6, 12) になる
		fmt.Println(p)
		scaleP(10) // p は (60, 120) になる
		fmt.Println(p)

		// レシーバーに対してメソッド呼び出ししたいときに役立つ
		example1()
	}

	// メソッド式(method expression)
	{
		p := Point{1, 2}
		q := Point{4, 6}

		distance := Point.Distance   // メソッド式
		fmt.Println(distance(p, q))  // 5
		fmt.Printf("%T\n", distance) // func(Point, Point) float64

		scale := (*Point).ScaleBy
		scale(&p, 2)
		fmt.Println(p)            // {2 4}
		fmt.Printf("%T\n", scale) // func(*Point, float64)

		path := Path{Point{1, 2}, Point{3, 4}}
		path.TranslateBy(Point{1, 2}, true)
		fmt.Println(path)
	}
}

type Rocket struct{}

func (r *Rocket) Launch() { fmt.Println("Launch!") }

func example1() {
	r := new(Rocket)
	time.AfterFunc(10*time.Second, func() { r.Launch() })
	// メソッド値なら、以下のように短くできる
	time.AfterFunc(10*time.Second, r.Launch)
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

// メソッド式を使って実行する処理を決める
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// path[i].Add(offset) か path[i].Sub(offset) のどちらかを呼び出す
		path[i] = op(path[i], offset)
	}
}
