package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	fmt.Println("Hello123")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){
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
