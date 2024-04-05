package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func get_port() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func main() {
	router := gin.Default()

	router.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	port := get_port()
	router.Run(port)
}
