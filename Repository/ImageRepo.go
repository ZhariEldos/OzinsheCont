package Repository

import (
	"fmt"
	"os"
)

type ImageRepository struct {
	filepath string
}

func NewImageRepository(filepath string) *ImageRepository {
	return &ImageRepository{filepath: filepath}
}

func (r *ImageRepository) FindThisImage(name string, folder string) ([]byte, error) {
	byteFile, err := os.ReadFile(fmt.Sprintf("%s/%s/%s", r.filepath, folder, name))
	if err != nil {
		return nil, err
	}
	return byteFile, nil
}
