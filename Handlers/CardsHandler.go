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
	if err != nil {
		c.JSON(http.StatusInternalServerError, Structs.NewApiError(err.Error(), "FindThisCard() {} (CardsRepo)"))
		return
	}
	c.JSON(http.StatusOK, card)
}
