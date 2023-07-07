package main

import (
	"dudan/database"
	"dudan/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	database.DbConnect()
	router := gin.Default()

	// Apply CORS middleware to all routes
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	router.POST("/registeruser", user.RegisterUser)
	router.GET("/loginuser", user.LoginUser)
	router.GET("/fetchusers", user.FetchUsers)
	router.Run("localhost:8081")
}
