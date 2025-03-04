package Methods

import (
	"fmt"
	"ozinsheproject/Handlers"

	"github.com/gin-gonic/gin"
)

func InitCardsMethods(r *gin.Engine, cardsHandler *Handlers.CardsHandler) {
	fmt.Println("\nCards Handlers:\n ")
	r.GET("/Card", cardsHandler.FindAllCards)
	r.GET("/Card/:id", cardsHandler.FindThisCard)
	r.POST("/Card", cardsHandler.CreateCard)
	r.PUT("/Card/:id", cardsHandler.UpdateCard)
	r.DELETE("/Card/:id", cardsHandler.DeleteCards)
}
