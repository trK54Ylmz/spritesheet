package util

import "image"

func CheckSizes(images []*image.Image) bool {
	w, h := 0, 0

	for index := range images {
		point := (*images[index]).Bounds().Size()

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
