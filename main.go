package main

import (
	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/api/url"
	"github.com/devchrischen/url-shortener/config"
	"github.com/devchrischen/url-shortener/lib/db"
	"github.com/devchrischen/url-shortener/lib/redis"
)

func main() {
	config.Init()
	db.Init()
	redis.Init()

	router := gin.Default()

	api := router.Group("")

	url.Route(api)

	router.Run()
}
