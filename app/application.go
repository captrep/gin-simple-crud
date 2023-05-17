package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	router.Run(":8080")
}
