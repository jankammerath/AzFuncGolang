package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
)

func get_port() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func getCPUName() string {
	info, err := cpu.Info()
	if err != nil {
		return err.Error()
	}
	if len(info) > 0 {
		return info[0].ModelName
	}
	return ""
}

func main() {
	router := gin.Default()

	router.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/api/cpuinfo", func(c *gin.Context) {
		cpuName := getCPUName()
		c.JSON(http.StatusOK, gin.H{
			"cpu": cpuName,
		})
	})

	port := get_port()
	router.Run(port)
}
