package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type params struct {
	x     int
	y     int
	scale float64
}

func (p *params) set(values url.Values) {
	x, err := strconv.Atoi(values.Get("x"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse x: %v\n", err)
	}
	y, err := strconv.Atoi(values.Get("y"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse y: %v\n", err)
	}
	scale, err := strconv.ParseFloat(values.Get("scale"), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse scale: %v\n", err)
	}
	p.x = x
	p.y = y
	p.scale = scale
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var p params
	p.set(r.URL.Query())
	render(w, &p)
}

func render(out io.Writer, p *params) {
	const (
		width, height = 1024, 1024
	)
	xmin := -2 / p.scale
	ymin := -2 / p.scale
	xmax := +2 / p.scale
	ymax := +2 / p.scale

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py-p.y)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px-p.x)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
