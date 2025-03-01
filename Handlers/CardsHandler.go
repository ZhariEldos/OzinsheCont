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
