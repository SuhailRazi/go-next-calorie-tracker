package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")

	// Initializing port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("entry/create")
	router.GET("/entries")
	router.GET("/entries/:id")
	router.GET("/ingredient/:ingredient")

	router.PUT("/entry/update/:id")
	router.PUT("/ingredient/update/:id")
	router.DELETE("/entry/delete/:id")

	router.Run(":" + port)

}
