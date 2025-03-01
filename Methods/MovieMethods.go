package Methods

import (
	"ozinsheproject/Handlers"

	"github.com/gin-gonic/gin"
)

func InitMovieMethods(r *gin.Engine, movieHandler *Handlers.MovieHandler) {
	r.GET("/Movie", movieHandler.FindAllMovie)
	r.GET("/Movie/:id", movieHandler.FindThisMovie)
	r.PUT("/Movie/:id", movieHandler.UpdateMovie)
	r.DELETE("/Movie/:id", movieHandler.DeleteMovie)
	r.POST("/Movie", movieHandler.CreateMovie)
}
