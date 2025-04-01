package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/play-area-workspace/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	fmt.Printf("Hello123 %s", os.Getenv("PORT"))

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:id", func(c *gin.Context) {
		userID := c.Param("id") // Extracts `id` from URL path

		// JSON response
		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
			"message": "User found",
		})
	})

	router.Run()
}
