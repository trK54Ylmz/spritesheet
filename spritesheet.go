package main

import (
	"image"

	"github.com/trk54ylmz/spritesheet/pkg/folder"
	"github.com/trk54ylmz/spritesheet/pkg/picture"
)

// Process the input images and generate the output image
func Process(input, output *string) error {
	fr := folder.InitFileReader(input)

	entries, err := fr.ListDir()
	if err != nil {
		return err
	}

	images := make([]*image.Image, len(entries))
	for index := range entries {
		f, err := fr.Read(entries[index])
		if err != nil {
			return err
		}

		defer f.Close()

		ir := picture.InitImageReader(f)

		image, err := ir.Read()
		if err != nil {
			return err
		}

		images[index] = image
	}

	it := picture.InitImageTrim(images)

	trimmed, width, height, err := it.Trim()
	if err != nil {
		return err
	}

	iw := picture.NewImageWriter(*width, *height, len(trimmed))

	for index := range trimmed {
		iw.Append(index, trimmed[index])
	}

	return iw.Write(*output)
}
