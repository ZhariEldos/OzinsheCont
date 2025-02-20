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

func (h *CategoryHandler) FindCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "FindCategoryByID() {} (CategoryHandler)"))
		return
	}
	category, err := h.CategoryRepo.FindCategoryByID(c, id)
	if err != nil && err.Error() == "no rows in result set" {
		c.JSON(http.StatusNotFound, Structs.NewApiError("Category with this id not found", "FindCategoryByID() {} (CategoryRepo)"))
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindCategoryByID() {} (CategoryRepo)"))
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) FindAllCategories(c *gin.Context) {
	categories, err := h.CategoryRepo.FindAllCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindCategoryByID() {} (CategoryRepo)"))
		return
	}
	c.JSON(http.StatusOK, categories)
}
