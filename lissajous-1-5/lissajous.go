package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette []color.Color

const (
	whiteIndex = 0 // パレットの最初の色
	blackIndex = 1 // パレットの次の色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	generatePalette()
	lissajous(os.Stdout)
}

func generatePalette() {
	palette = append(palette, color.White)
	for i := 1; i <= 10; i++ {
		r := uint8(rand.Intn(256))
		g := uint8(rand.Intn(256))
		b := uint8(rand.Intn(256))
		palette = append(palette, color.RGBA{r, g, b, 1})
	}
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 発振器 x が完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンパスは [-size..+size] の範囲を扱う
		nframes = 64    // アニメーションのフレーム数
		delay   = 8     // 10ms 単位でのフレーム間の遅延
	)
	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			paletteNo := uint8(rand.Intn(len(palette)))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), paletteNo)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}
