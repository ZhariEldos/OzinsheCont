package main

import (
	"context"
	"ozinsheproject/Handlers"
	"ozinsheproject/Repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	r := gin.Default()

	conn, err := connectToDB()
	if err != nil {
		panic(err)
	}
	cardsRepo := Repository.NewCardsRepository(conn)
	movieRepo := Repository.NewMovieRepository(conn)
	categoryRepo := Repository.NewCategoryRepository(conn)
	cardsHandler := Handlers.NewCardsHandler(cardsRepo)
	movieHandler := Handlers.NewMovieHandler(movieRepo)
	categoryHandler := Handlers.NewCategoryHandler(categoryRepo)

	// TODO: All http methods devide to another file
	// Movie:
	r.GET("/Movie", movieHandler.FindAllMovie)
	r.GET("/Movie/:id", movieHandler.FindThisMovie)
	r.PUT("/Movie/:id", movieHandler.UpdateMovie)
	r.DELETE("/Movie/:id", movieHandler.DeleteMovie)
	r.POST("/Movie", movieHandler.CreateMovie)
	// Category:
	r.GET("/Category/:id", categoryHandler.FindThisCategory)
	r.GET("/Category", categoryHandler.FindAllCategories)
	r.POST("/Category", categoryHandler.CreateCategory)
	r.PUT("/Category/:id", categoryHandler.UpdateCategory)
	r.DELETE("/Category/:id", categoryHandler.DeleteCategory)
	// Cards:
	r.GET("/Card", cardsHandler.FindAllCards)
	r.GET("/Card/:id", cardsHandler.FindThisCard)

	r.Run()
}

func connectToDB() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), "postgresql://postgres:100100100AaAa@localhost:5432/ozinshe")
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
