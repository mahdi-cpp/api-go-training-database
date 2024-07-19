package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	router = gin.Default()
)

// Run will instagram the server
func Run() {

	err := router.SetTrustedProxies([]string{"192.168.1.2"})
	if err != nil {
		return
	}

	router.Use(CORSMiddleware())

	// Define the file system for serving static files
	fs := http.Dir("static")

	//Serve static HTML files from the "static" directory
	router.StaticFS("/v2/ali", fs)

	getRoutes()
	err = router.Run(":8095")
	if err != nil {
		return
	}
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v2 := router.Group("/v2")
	addApiRoutes(v2)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
