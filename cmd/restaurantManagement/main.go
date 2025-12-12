package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suryanshvermaa/restaurantManagement/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	// all middlewares and routes
	router.Use(gin.Logger())

	// health route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"success": true,
			"message": "server is healthy",
		})
	})
	router.Run(":" + port)
}
