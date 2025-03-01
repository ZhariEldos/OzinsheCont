package Methods

import (
	"ozinsheproject/Handlers"

	"github.com/gin-gonic/gin"
)

func InitCardsMethods(r *gin.Engine, cardsHandler *Handlers.CardsHandler) {
	r.GET("/Card", cardsHandler.FindAllCards)
	r.GET("/Card/:id", cardsHandler.FindThisCard)
	r.POST("/Card", cardsHandler.CreateCard)
}
