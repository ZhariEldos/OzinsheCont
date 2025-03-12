package Repository

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (r *ImageRepository) CreateImage(c *gin.Context, image *multipart.FileHeader, folder string) (name string, err error) {
	filename := fmt.Sprintf("%s%s", uuid.NewString(), filepath.Ext(image.Filename))
	filepath := fmt.Sprintf("%s/%s/%s", r.filepath, folder, filename)
	err = c.SaveUploadedFile(image, filepath)
	if err != nil {
		return "", err
	}
	return filename, nil
}
