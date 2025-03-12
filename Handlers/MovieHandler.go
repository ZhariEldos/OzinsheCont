package Handlers

import (
	"net/http"
	"ozinsheproject/Repository"
	"ozinsheproject/Structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	MovieRepo *Repository.MovieRepository
}

func NewMovieHandler(MovieRepo *Repository.MovieRepository) *MovieHandler {
	return &MovieHandler{MovieRepo: MovieRepo}
}

type requestForCreateMovie struct {
	MovieTitle  string
	Director    string
	Producer    string
	Description string
	Realesed    int
	Category    []Structs.Category
	Cards       []Structs.Cards
	URLPoster   string
}
type requestForUpdateMovie struct {
	MovieTitle  string
	Director    string
	Producer    string
	Description string
	Realesed    int
	Category    []Structs.Category
	Cards       []Structs.Cards
	URLPoster   string
}

// TODO: Make search for movie by params
func (h *MovieHandler) FindAllMovie(c *gin.Context) {
	Movies, err := h.MovieRepo.FindAllMovies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindAllMovie() {} (MovieRepo)"))
		return
	}
	c.JSON(http.StatusOK, Movies)
}

func (h *MovieHandler) FindThisMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "FindThisMovie() {} (MovieHandler)"))
		return
	}
	Movie, err := h.MovieRepo.FindThisMovie(c, id)
	if Movie.ID == -1 {
		c.JSON(http.StatusNotFound, Structs.NewApiError("Category with this id not found", "FindThisMovie() {} (MovieRepo)"))
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindThisMovie() {} (MovieRepo)"))
		return
	}
	c.JSON(http.StatusOK, Movie)
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var request requestForCreateMovie
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError(err.Error(), "CreateMovie() {} (MovieHandler)"))
	}
	movie := Structs.Movie{
		MovieTitle:  request.MovieTitle,
		Director:    request.Director,
		Producer:    request.Producer,
		Description: request.Description,
		Realesed:    request.Realesed,
		Category:    request.Category,
		Cards:       request.Cards,
		URLPoster:   request.URLPoster,
	}
	id, err := h.MovieRepo.CreateMovie(c, movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError(err.Error(), "CreateMovie() {} (MovieRepo)"))
	}
	c.JSON(http.StatusOK, id)
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "UpdateMovie() {} (MovieHandler)"))
		return
	}
	var request requestForUpdateMovie
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError(err.Error(), "UpdateMovie() {} (MovieHandler)"))
	}
	movie := Structs.Movie{
		ID:          id,
		MovieTitle:  request.MovieTitle,
		Director:    request.Director,
		Producer:    request.Producer,
		Description: request.Description,
		Realesed:    request.Realesed,
		Category:    request.Category,
		Cards:       request.Cards,
		URLPoster:   request.URLPoster,
	}
	err = h.MovieRepo.UpdateMovie(c, movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "UpdateMovie() {} (MovieRepo)"))
	}
	c.Status(http.StatusOK)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "UpdateMovie() {} (MovieHandler)"))
	}
	err = h.MovieRepo.DeleteMovie(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "UpdateMovie() {} (MovieRepo)"))
	}
}
