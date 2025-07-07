package main

import (
	"github.com/a1l2v/REST_API_WITH_GORM.git/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "pong",
		})
	})
	router.Run()
}