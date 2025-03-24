package Methods

import (
	"fmt"
	"ozinsheproject/Handlers"

	"github.com/gin-gonic/gin"
)

func InitImageMethods(r *gin.Engine, imageHandler *Handlers.ImageHandler) {
	fmt.Println("\nImage Handlers:\n ")
	r.GET("/image/:folder/:name", imageHandler.FindThisImage)
	r.POST("/image/:folder", imageHandler.CreateImage)
}
