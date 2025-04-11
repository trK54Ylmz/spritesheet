package io

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"
)

type ImageWriter struct {
	image *image.RGBA

	width  int
	height int
	part   int
}

func NewImageWriter(width, height, part int) *ImageWriter {
	i := new(ImageWriter)
	i.width = width
	i.height = height
	i.part = part

	border := image.Rectangle{image.Point{0, 0}, image.Point{width * part, height}}
	i.image = image.NewRGBA(border)

	return i
}

func (w *ImageWriter) Append(index int, image *image.Image) {
	img := (*image)

	for i := 0; i < w.width; i++ {
		x := (w.width * index) + i
		for y := 0; y < w.height; y++ {
			r, g, b, a := img.At(i, y).RGBA()
			w.image.SetRGBA(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
}

func (w *ImageWriter) Write(path string) error {
	var b bytes.Buffer

	png.Encode(&b, w.image)

	return os.WriteFile(path, b.Bytes(), 0644)
}
