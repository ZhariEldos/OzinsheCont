package Handlers

import (
	"fmt"
	"net/http"
	"ozinsheproject/Repository"
	"ozinsheproject/Structs"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	ImageRepo *Repository.ImageRepository
}

func NewImageHandler(ImageRepo *Repository.ImageRepository) *ImageHandler {
	return &ImageHandler{ImageRepo: ImageRepo}
}

func (h *ImageHandler) FindThisImage(c *gin.Context) {
	name := c.Param("name")
	folder := c.Param("folder")
	fileName := filepath.Base(name)
	image, err := h.ImageRepo.FindThisImage(name, folder)
	if err != nil {
		c.JSON(http.StatusNotFound, Structs.NewApiError(err.Error(), "FindThisImage() {} (ImageRepo)"))
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Data(http.StatusOK, "application/octet-stream", image)
}
