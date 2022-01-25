package main

import (
	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/api/url"
	"github.com/devchrischen/url-shortener/lib/db"
)

func main() {
	router := gin.Default()
	db.Init()

	api := router.Group("")

	url.Route(api)

	router.Run()
}
