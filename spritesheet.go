package main

import (
	"image"
	"log"

	"github.com/trk54ylmz/spritesheet/pkg/io"
	"github.com/trk54ylmz/spritesheet/pkg/picture"
	"github.com/trk54ylmz/spritesheet/util"
)

// Process the input images and generate the output image
func Process(input, output *string, trim *bool) error {
	fr := io.NewFileReader(input)

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

		ir := picture.NewImageReader(f)

		image, err := ir.Read()
		if err != nil {
			return err
		}

		defer ir.Close()

		images[index] = image
	}

	log.Println("Checking file sizes ...")

	identical := util.CheckSizes(images)
	if !identical {
		return util.ErrSizeNotIdentical
	}

	var width *int
	var height *int
	if *trim {
		it := picture.NewImageTrim(images)

		images, width, height, err = it.Trim()
		if err != nil {
			return err
		}
	} else {
		size := (*images[0]).Bounds().Size()

		width = &size.X
		height = &size.Y
	}

	iw := io.NewImageWriter(*width, *height, len(images))

	for index := range images {
		iw.Append(index, images[index])
	}

	return iw.Write(*output)
}
