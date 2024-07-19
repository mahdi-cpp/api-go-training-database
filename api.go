package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-gallery/repository"
)

func addApiRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/music")

	// Define your API endpoints
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	router.GET("/users", func(c *gin.Context) {
		c.JSON(210, gin.H{"users": repository.GetUsers()})
	})

	router.GET("/posts", func(c *gin.Context) {
		c.JSON(210, gin.H{"posts": repository.GetPosts()})
	})
}
