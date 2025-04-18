package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(cfg *Config) *gin.Engine {
	router := gin.Default()
	// Register routes
	router.GET("/genPlist", GeneratePlist)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "plistserver is running on Vercel. See https://github.com/missuo/vercel-plistserver for more information. \n\n Made with ❤️ by Vincent."})
	})

	// Catch all other routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	return router
}
