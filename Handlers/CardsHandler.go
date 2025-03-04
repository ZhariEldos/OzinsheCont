package Handlers

import (
	"net/http"
	"ozinsheproject/Repository"
	"ozinsheproject/Structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CardsHandler struct {
	CardsRepo *Repository.CardsRepository
}

func NewCardsHandler(CardsRepo *Repository.CardsRepository) *CardsHandler {
	return &CardsHandler{CardsRepo: CardsRepo}
}

type requestForCreateCard struct {
	CardsTitle string
	URLPicture string // TODO: Pictures system
}
type requestForUpdateCard struct {
	CardsTitle string
	URLPicture string // TODO: Pictures system
}

func (h *CardsHandler) FindAllCards(c *gin.Context) {
	cards, err := h.CardsRepo.FindAllCards(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindAllCards() {} (CardsRepo)"))
		return
	}
	c.JSON(http.StatusOK, cards)
}

func (h *CardsHandler) FindThisCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "FindThisCard() {} (CardsHandler)"))
		return
	}
	card, err := h.CardsRepo.FindThisCard(c, id)
	if (err != nil) && (card == Structs.Cards{}) {
		c.JSON(http.StatusNotFound, Structs.NewApiError("card with this id not found", "FindThisCard() {} (CardsRepo)"))
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindThisCard() {} (CardsRepo)"))
		return
	}
	c.JSON(http.StatusOK, card)
}

func (h *CardsHandler) CreateCard(c *gin.Context) {
	var request requestForCreateCard
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError(err.Error(), "CreateCard() {} (CardsHandler)"))
		return
	}
	card := Structs.Cards{ID: 0, CardsTitle: request.CardsTitle, URLPicture: request.URLPicture}
	id, err := h.CardsRepo.CreateCard(c, card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "CreateCard() {} (CardsRepo)"))
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *CardsHandler) UpdateCard(c *gin.Context) {
	var request requestForUpdateCard
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError(err.Error(), "UpdateCard() {} (CardsHandler)"))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "UpdateCard() {} (CardsHandler)"))
		return
	}
	card := Structs.Cards{
		ID:         id,
		CardsTitle: request.CardsTitle,
		URLPicture: request.URLPicture,
	}
	err = h.CardsRepo.UpdateCard(c, card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "UpdateCard() {} (CardsRepo)"))
		return
	}
	c.Status(http.StatusOK)
}

func (h *CardsHandler) DeleteCards(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Structs.NewApiError("Id of movie have strange value or type of value", "DeleteCards() {} (CardsHandler)"))
		return
	}
	err = h.CardsRepo.DeleteCards(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "DeleteCards() {} (CardsRepo)"))
		return
	}
	c.Status(http.StatusOK)
}
