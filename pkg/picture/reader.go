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

// InitImageReader
func InitImageReader(file *os.File) *ImageReader {
	r := new(ImageReader)
	r.file = file

	return r
}

func (i *ImageReader) Read() (*image.Image, error) {
	img, _, err := image.Decode(i.file)
	if err != nil {
		return nil, err
	}

	return &img, nil
}
