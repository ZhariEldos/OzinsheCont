package Methods

import (
	"fmt"
	"ozinsheproject/Handlers"

	"github.com/gin-gonic/gin"
)

func InitMovieMethods(r *gin.Engine, movieHandler *Handlers.MovieHandler) {
	fmt.Println("\nMovies Handlers:\n ")
	r.GET("/Movie", movieHandler.FindAllMovie)
	r.GET("/Movie/:id", movieHandler.FindMovieByID)
	r.GET("/Movie/search", movieHandler.FindMovieByParams)
	r.POST("/Movie", movieHandler.CreateMovie)
	r.PUT("/Movie/:id", movieHandler.UpdateMovie)
	r.DELETE("/Movie/:id", movieHandler.DeleteMovie)
}
