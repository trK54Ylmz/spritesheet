package picture

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/trk54ylmz/spritesheet/pkg/util"
)

type ImageTrim struct {
	images []*image.Image
}

func InitImageTrim(images []*image.Image) *ImageTrim {
	i := new(ImageTrim)
	i.images = images

	return i
}

func (i *ImageTrim) size(image *image.Image) image.Point {
	img := *(image)

	return img.Bounds().Size()
}

func (i *ImageTrim) rgba(image *image.Image, x, y int) *color.RGBA {
	r, g, b, a := (*image).At(x, y).RGBA()

	return &color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

func (i *ImageTrim) CheckSize() bool {
	w, h := 0, 0

	for index := range i.images {
		point := i.size(i.images[index])

		if index == 0 {
			w = point.X
			h = point.Y
		} else {
			if w != point.X {
				return false
			}

			if h != point.Y {
				return false
			}
		}
	}

	return true
}

func (i *ImageTrim) OptimalSize(image *image.Image) (int, int, int, int) {
	size := i.size(image)

	l, t, r, b := 0, 0, size.X, size.Y

	found := false
	for i := 0; i < size.X; i++ {
		if found {
			break
		}

		for j := 0; j < size.Y; j++ {
			_, _, _, alpha := (*image).At(i, j).RGBA()

			if alpha != 0 {
				l = i
				found = true
				break
			}
		}
	}

	found = false
	for i := 0; i < size.Y; i++ {
		if found {
			break
		}

		for j := 0; j < size.X; j++ {
			_, _, _, alpha := (*image).At(j, i).RGBA()

			if alpha != 0 {
				t = i
				found = true
				break
			}
		}
	}

	found = false
	for i := size.X; i > l; i-- {
		if found {
			break
		}

		for j := 0; j < size.Y; j++ {
			_, _, _, alpha := (*image).At(i, j).RGBA()

			if alpha != 0 {
				r = i
				found = true
				break
			}
		}
	}

	found = false
	for i := size.Y; i > t; i-- {
		if found {
			break
		}

		for j := 0; j < size.X; j++ {
			_, _, _, alpha := (*image).At(j, i).RGBA()

			if alpha != 0 {
				b = i
				found = true
				break
			}
		}
	}

	return l, t, r, b
}

func (i *ImageTrim) Trim() ([]*image.Image, *int, *int, error) {
	log.Println("Checking file sizes ...")

	identical := i.CheckSize()
	if !identical {
		return nil, nil, nil, util.ErrSizeNotIdentical
	}

	// Get size of the pictures
	point := i.size(i.images[0])

	log.Printf("The original file size is %dx%d px\n", point.X, point.Y)

	l, t, r, b := math.MaxInt32, math.MaxInt32, 0, 0
	for index := range i.images {
		ll, lt, lr, lb := i.OptimalSize(i.images[index])
		if ll < l {
			l = ll
		}
		if lt < t {
			t = lt
		}
		if lr > r {
			r = lr
		}
		if lb > b {
			b = lb
		}

	}

	nw := point.X - l - (point.X - r)
	nh := point.Y - t - (point.Y - b)
	np := image.Rectangle{image.Point{0, 0}, image.Point{nw, nh}}

	log.Printf("The trimmed file size is %dx%d px\n", nw, nh)

	images := make([]*image.Image, len(i.images))

	for index := range i.images {
		rgba := image.NewRGBA(np)

		for x := 0; x < nw; x++ {
			for y := 0; y < nh; y++ {
				rgba.SetRGBA(x, y, *i.rgba(i.images[index], x+l, y+t))
			}
		}

		var img image.Image = rgba

		images[index] = &img
	}

	log.Printf("The final file size is %dx%d px\n", nw*len(images), nh)

	return images, &nw, &nh, nil
}
