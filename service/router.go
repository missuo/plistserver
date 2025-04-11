package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(cfg *Config) *gin.Engine {
	router := gin.Default()

	// Register routes
	router.GET("/genPlist", GeneratePlist)

	// Catch all other routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	return router
}
