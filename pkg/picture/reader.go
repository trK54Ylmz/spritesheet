package picture

import (
	"image"
	_ "image/png"
	"os"
)

// ImageReader is the read and decode images
type ImageReader struct {
	file *os.File
}

// NewImageReader creates a image reader for a file
func NewImageReader(file *os.File) *ImageReader {
	r := new(ImageReader)
	r.file = file

	return r
}

// Read content of the png file and decode
func (i *ImageReader) Read() (*image.Image, error) {
	img, _, err := image.Decode(i.file)
	if err != nil {
		return nil, err
	}

	return &img, nil
}
