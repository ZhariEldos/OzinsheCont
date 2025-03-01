package main

import (
	"context"
	"ozinsheproject/Handlers"
	"ozinsheproject/Methods"
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

	// TODO: Make a docs
	// TODO: All http methods devide to another file
	// Movie:
	Methods.InitMovieMethods(r, movieHandler)
	// Category:
	Methods.InitCategoryMethods(r, categoryHandler)
	// Cards:
	Methods.InitCardsMethods(r, cardsHandler)

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
