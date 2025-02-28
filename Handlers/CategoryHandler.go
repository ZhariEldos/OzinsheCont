package Handlers

import (
	"net/http"
	"ozinsheproject/Repository"
	"ozinsheproject/Structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryRepo *Repository.CategoryRepository
}

func NewCategoryHandler(CategoryRepo *Repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{CategoryRepo: CategoryRepo}
}

type requestForCreateCategory struct {
	CategoryTitle string
}

type requestForUpdateCategory struct {
	CategoryTitle string
}

func (h *CategoryHandler) FindThisCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "FindThisCategory() {} (CategoryHandler)"))
		return
	}
	category, err := h.CategoryRepo.FindThisCategory(c, id)
	if (err != nil) && (category == Structs.Category{}) {
		c.JSON(http.StatusNotFound, Structs.NewApiError("Category with this id not found", "FindThisCategory() {} (CategoryRepo)"))
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindThisCategory() {} (CategoryRepo)"))
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) FindAllCategories(c *gin.Context) {
	categories, err := h.CategoryRepo.FindAllCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindAllCategories() {} (CategoryRepo)"))
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var request requestForCreateCategory
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError(err.Error(), "CreateCategory() {} (CategoryHandler)"))
		return
	}
	category := Structs.Category{ID: 0, CategoryTitle: request.CategoryTitle}
	id, err := h.CategoryRepo.CreateCategory(c, category)
	if (err != nil) || (id == -1) {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "CreateCategory() {} (CategoryRepo)"))
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "UpdateCategory() {} (CategoryHandler)"))
		return
	}
	var request requestForUpdateCategory
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError(err.Error(), "UpdateCategory() {} (CategoryHandler)"))
	}
	category := Structs.Category{ID: id, CategoryTitle: request.CategoryTitle}
	err = h.CategoryRepo.UpdateCategory(c, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "UpdateCategory() {} (CategoryRepo)"))
		return
	}
	c.Status(http.StatusOK)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "DeleteCategory() {} (CategoryHandler)"))
		return
	}
	err = h.CategoryRepo.DeleteCategory(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "DeleteCategory() {} (CategoryRepo)"))
	}
	c.Status(http.StatusOK)
}
