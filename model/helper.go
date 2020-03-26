package model

import (
	"github.com/labstack/gommon/random"
	"io"
	"mime/multipart"
	"os"
	"time"
)

func InsertImage(header *multipart.FileHeader, dest string) (string, error) {
	src, err := header.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	randoms := random.New()

	filename := time.Now().Format(time.RFC3339Nano) + randoms.String(10, random.Alphanumeric) + header.Filename

	dst, err := os.Create(dest + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return filename, nil
}
