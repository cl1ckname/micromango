package utils

import (
	"io"
	"mime/multipart"
)

func ReadFormFile(formFile *multipart.FileHeader) ([]byte, error) {
	file, err := formFile.Open()
	if err != nil {
		return nil, err
	}
	imageBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return imageBytes, nil
}
