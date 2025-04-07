package io

import (
	"log"
	"os"
	"strings"

	"github.com/trk54ylmz/spritesheet/pkg/util"
)

type FileReader struct {
	path *string
}

// NewFileReader creates a file reader for a path
func NewFileReader(path *string) *FileReader {
	f := new(FileReader)
	f.path = path

	return f
}

// ListDir for a path
func (f *FileReader) ListDir() ([]string, error) {
	if _, err := os.Stat(*f.path); err != nil {
		if os.IsNotExist(err) {
			return nil, util.ErrFolderNotExist
		}

		return nil, err
	}

	entries, err := os.ReadDir(*f.path)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.HasPrefix(entry.Name(), ".DS_Store") {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".png") {
			log.Println("There is a non png file")
			continue
		}

		files = append(files, strings.Join([]string{*f.path, entry.Name()}, string(os.PathSeparator)))
	}

	return files, nil
}

// Read content of the given png file
func (f *FileReader) Read(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
