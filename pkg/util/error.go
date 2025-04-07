package util

import "errors"

var (
	ErrSizeNotIdentical = errors.New("image sizes are not identical")
	ErrFolderNotExist   = errors.New("input folder does not exist")
)
