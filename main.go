package main

import (
	"context"
	"fmt"
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
	imageRepo := Repository.NewImageRepository("image")
	cardsHandler := Handlers.NewCardsHandler(cardsRepo)
	movieHandler := Handlers.NewMovieHandler(movieRepo)
	categoryHandler := Handlers.NewCategoryHandler(categoryRepo)
	imageHandler := Handlers.NewImageHandler(imageRepo)

	// TODO: Make a docs
	// Movie:
	Methods.InitMovieMethods(r, movieHandler)
	// Category:
	Methods.InitCategoryMethods(r, categoryHandler)
	// Cards:
	Methods.InitCardsMethods(r, cardsHandler)
	// Image:
	Methods.InitImageMethods(r, imageHandler)

	fmt.Println()

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
