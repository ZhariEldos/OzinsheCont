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

// @Summary		Find image
// @Description Find the image in server
// @Param		folder path string true "Path to picture"
// @Param		name path string true "Picture name"
// @Tags		Images
// @Produce		png
// @Success		200 {object} []byte "PosterPNG, PosterJP(E)G"
// @Failure		404 {object} Structs.ApiError
// @Router		/image/{folder}/{name} [get]
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

// @Summary		Create image
// @Description Put image to server
// @Param		folder path string true "Path to picture"
// @Tags		Images
// @Accept		mpfd
// @Produce		json
// @Success		200 {object} string "Name of new image"
// @Failure		500 {object} Structs.ApiError
// @Failure		400 {object} Structs.ApiError
// @Router		/image/{folder} [post]
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
