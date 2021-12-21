package fs

import (
	"backendOneLessons/lesson4/internal/pkg/image"
	"context"
	"os"
)

type fs struct {
	basePath string
}

func (f fs) Save(ctx context.Context, filename string, data []byte) (string, error) {
	path := f.basePath + filename
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return "", err
	}

	return path, nil
}

func New(path string) image.Storage {
	return fs{
		basePath: path,
	}
}
