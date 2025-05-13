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

// @Summary		Find category by ID
// @Description Find category from DataBase by ID
// @Tags		Categories
// @Param		id path int true "Category ID"
// @Produce		json
// @Success		200 {object} Structs.Category
// @Failure		500 {object} Structs.ApiError
// @Failure		400 {object} Structs.ApiError
// @Failure		404 {object} Structs.ApiError
// @Router		/Category/{id} [get]
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

// @Summary		Find all categories
// @Description Find all categories from DataBase by ID
// @Tags		Categories
// @Produce		json
// @Success		200 {object} []Structs.Category
// @Failure		500 {object} Structs.ApiError
// @Router		/Category [get]
func (h *CategoryHandler) FindAllCategories(c *gin.Context) {
	categories, err := h.CategoryRepo.FindAllCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindAllCategories() {} (CategoryRepo)"))
		return
	}
	c.JSON(http.StatusOK, categories)
}

// @Summary		Create a category
// @Description Create new category and put it to database
// @Tags		Categories
// @Accept		json
// @Produce		json
// @Success		200 {object} int "ID new category"
// @Failure		500 {object} Structs.ApiError
// @Failure		400 {object} Structs.ApiError
// @Router		/Category [post]
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

// @Summary		Update category info
// @Description Update category information from database
// @Tags		Categories
// @Param		id path int true "Category ID"
// @Accept		json
// @Produce		json
// @Success		200
// @Failure		500 {object} Structs.ApiError
// @Failure		400 {object} Structs.ApiError
// @Router		/Category/{id} [put]
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

// @Summary		Delete category
// @Description Delete category from database
// @Tags		Categories
// @Param		id path int true "Category ID"
// @Produce		json
// @Success		200
// @Failure		500 {object} Structs.ApiError
// @Failure		400 {object} Structs.ApiError
// @Router		/Category/{id} [delete]
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
