package main

import (
	"context"
	"fmt"
	"ozinsheproject/Handlers"
	"ozinsheproject/Methods"
	"ozinsheproject/Repository"

	"ozinsheproject/docs"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Ozinshe API
// @version         1.0
// @description     API for search movies by tags and publish them

// @host      localhost:8080
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

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println()

	r.Run()
}

func connectToDB() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), "postgresql://ozinshe_user:sc4hGncGEXhWiHhcXJmtC4b8wPotcOIY@dpg-d2jklf3ipnbc73ba0icg-a/ozinshe")
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
