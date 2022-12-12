package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
)

const (
	width  = 600
	height = 600
)

func main() {
	vs := makeValues()
	createImage(os.Stdout, vs)
}
func makeValues() [][]uint8 {
	vs := make([][]uint8, 0, width)
	for i := 0; i < width; i++ {
		r := make([]uint8, 0, height)
		for j := 0; j < height; j++ {
			v := uint8(255 * i / width)
			r = append(r, v)
		}
		vs = append(vs, r)
	}
	return vs
}

func makeValueWithI(ix int) [][]uint8 {
	vs := make([][]uint8, 0, width)
	for i := 0; i < width; i++ {
		r := make([]uint8, 0, height)
		for j := 0; j < height; j++ {
			v := uint8(255 * i / width * ix / 100)
			r = append(r, v)
		}
		vs = append(vs, r)
	}
	return vs
}

func makeNewImage(values [][]uint8) *image.Paletted {
	palette := make([]color.Color, 256)

	for i := 0; i < 256; i++ {
		palette[i] = color.Gray{255 - uint8(i)}
	}

	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			img.SetColorIndex(x, y, uint8(values[y][x]))
		}
	}
	return img
}

func createImage(w io.Writer, values [][]uint8) {

	imgs := []*image.Paletted{}

	for i := 0; i <= 100; i += 10 {
		vs := makeValueWithI(i)
		img := makeNewImage(vs)
		imgs = append(imgs, img)
	}

	anim := gif.GIF{Delay: make([]int, len(imgs)), Image: imgs}
	if err := gif.EncodeAll(w, &anim); err != nil {
		panic(err)
	}
	// img := makeNewImage()
	// for y := 0; y < int(height); y++ {
	// 	for x := 0; x < int(width); x++ {
	// 		img.SetColorIndex(x, y, uint8(values[y][x]))
	// 	}
	// }

}
