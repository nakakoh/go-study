package geometry

import (
	"math"
)

type Point struct{ X, Y float64 }

// 昔ながらの関数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 同じだが、Point型のメソッドとして
// p はメソッドのレシーバー (receiver)
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path は点を直線で結びつける道のりです。
type Path []Point

// Distance は Path に沿って進んだ距離を返します。
// レシーバーが違えばメソッド名は同じものが使える
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
