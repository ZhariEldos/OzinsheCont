package Methods

import (
	"fmt"
	"ozinsheproject/Handlers"

	"github.com/gin-gonic/gin"
)

func InitCategoryMethods(r *gin.Engine, categoryHandler *Handlers.CategoryHandler) {
	fmt.Println("\nCategories Handlers:\n ")
	r.GET("/Category/:id", categoryHandler.FindThisCategory)
	r.GET("/Category", categoryHandler.FindAllCategories)
	r.POST("/Category", categoryHandler.CreateCategory)
	r.PUT("/Category/:id", categoryHandler.UpdateCategory)
	r.DELETE("/Category/:id", categoryHandler.DeleteCategory)
}
