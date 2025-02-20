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
	movieRepo := Repository.NewMovieRepository(conn)
	categoryRepo := Repository.NewCategoryRepository(conn)
	movieHandler := Handlers.NewMovieHandler(movieRepo)
	categoryHandler := Handlers.NewCategoryHandler(categoryRepo)

	r.GET("/Movie", movieHandler.FindAllMovie)
	r.GET("/Movie/:id", movieHandler.FindThisMovie)
	r.POST("/Movie", movieHandler.CreateMovie)
	r.GET("/Category/:id", categoryHandler.FindCategoryByID)
	r.GET("/Category", categoryHandler.FindAllCategories)
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
