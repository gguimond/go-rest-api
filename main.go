package main

import (
	"go-rest-api/controllers"
	"go-rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	database := client.Database("quickstart")
	r.Use(middleware.DB(database))

	// Release resource when the main
	// function is returned.
	defer close(client, ctx, cancel)

	r.GET("/", controllers.GetEpisodes)
	r.POST("/", controllers.InsertOneEpisode)
	r.PUT("/:id", controllers.UpdateOneEpisode)
	r.DELETE("/:id", controllers.DeleteOneEpisode)

	r.Run()
}
