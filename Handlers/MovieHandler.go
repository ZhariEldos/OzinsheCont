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

// @Summary		Show all movies
// @Description Show all movies in database and show their data
// @Tags		Movies
// @Produce		json
// @Success		200 {object} []Structs.Movie
// @Failure		500 {object} Structs.ApiError
// @Router		/Movie [get]
func (h *MovieHandler) FindAllMovie(c *gin.Context) {
	Movies, err := h.MovieRepo.FindAllMovies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindAllMovie() {} (MovieRepo)"))
		return
	}
	c.JSON(http.StatusOK, Movies)
}

// @Summary		Find movie by ID
// @Description Find movie from DataBase by ID
// @Tags		Movies
// @Params		id path int true "Movie ID"
// @Produce		json
// @Success		200 {object} Structs.Movie
// @Failure		500 {object} Structs.ApiError
// @Failure		400 {object} Structs.ApiError
// @Failure		404 {object} Structs.ApiError
// @Router		/Movie/{id} [get]
func (h *MovieHandler) FindMovieByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "FindMovieByID() {} (MovieHandler)"))
		return
	}
	Movie, err := h.MovieRepo.FindMovieByID(c, id)
	if Movie.ID == -1 {
		c.JSON(http.StatusNotFound, Structs.NewApiError("Category with this id not found", "FindMovieByID() {} (MovieRepo)"))
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindMovieByID() {} (MovieRepo)"))
		return
	}
	c.JSON(http.StatusOK, Movie)
}

// @Summary		(WIP) Search a movie
// @Description Search a movie by params, that you bind in json request
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Success		200 {object} Structs.Movie
// @Failure		500 {object} Structs.ApiError
// @Failure		400 {object} Structs.ApiError
// @Router		/Movie/search [get]
func (h *MovieHandler) FindMovieByParams(c *gin.Context) {
	var request Structs.FilmSearchParams
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Form of search is incorrect", "FindMovieByParams() {} (MovieHandler)"))
		return
	}
	movies, err := h.MovieRepo.FindMovieByParams(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movies)
}

// @Summary		Create a movie
// @Description Create new movie and put its to database
// @Tags		Movies
// @Accept		json
// @Produce		json
// @Success		200 {object} int
// @Failure		400 {object} Structs.ApiError
// @Router		/Movie [post]
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

// @Summary		Update existing movie
// @Description Update existing movie and put it to database
// @Tags		Movies
// @Params		id path int true "Movie ID"
// @Accept		json
// @Produce		json
// @Success		200
// @Failure		400 {object} Structs.ApiError
// @Failure		500 {object} Structs.ApiError
// @Router		/Movie/{id} [put]
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

// @Summary		Delete a movie
// @Description Delete existing movie from database
// @Tags		Movies
// @Params		id path int true "Movie ID"
// @Accept		json
// @Produce		json
// @Success		200
// @Failure		400 {object} Structs.ApiError
// @Failure		500 {object} Structs.ApiError
// @Router		/Movie/{id} [delete]
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
