package Methods

import (
	"ozinsheproject/Handlers"

	"github.com/gin-gonic/gin"
)

func InitCategoryMethods(r *gin.Engine, categoryHandler *Handlers.CategoryHandler) {
	r.GET("/Category/:id", categoryHandler.FindThisCategory)
	r.GET("/Category", categoryHandler.FindAllCategories)
	r.POST("/Category", categoryHandler.CreateCategory)
	r.PUT("/Category/:id", categoryHandler.UpdateCategory)
	r.DELETE("/Category/:id", categoryHandler.DeleteCategory)
}
