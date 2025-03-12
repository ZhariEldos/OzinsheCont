package Handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"ozinsheproject/Repository"
	"ozinsheproject/Structs"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	ImageRepo *Repository.ImageRepository
}

type createImageRequest struct {
	Image *multipart.FileHeader `form:"image"`
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

func (h *ImageHandler) CreateImage(c *gin.Context) {
	var request createImageRequest
	folder := c.Param("folder")
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError(err.Error(), "CreateImage() {} (ImageHandler)"))
		return
	}
	name, err := h.ImageRepo.CreateImage(c, request.Image, folder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "CreateImage() {} (ImageRepo)"))
		return
	}
	c.JSON(http.StatusOK, name)
}
