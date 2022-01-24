package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "short url created!")
	})

	r.GET("/:shortCode", func(c *gin.Context) {
		c.String(http.StatusOK, "successfully redirect!")
	})

	r.Run()
}
