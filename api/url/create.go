package url

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/entities/edb"
	"github.com/devchrischen/url-shortener/lib/apires"
	"github.com/devchrischen/url-shortener/lib/db"
	h "github.com/devchrischen/url-shortener/lib/hash"
	surl "github.com/devchrischen/url-shortener/services/url"
)

type shortenRequest struct {
	OriginalUrl string `json:"original_url" form:"original_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	// validate request
	var req shortenRequest
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
	}
	// check url duplicate => if exists in db, return the existing hash value
	urlService := surl.New(db.DB)
	hashValue, exist, err := urlService.CheckUrlExist(req.OriginalUrl)
	if err != nil {
		fmt.Println(err)
	}
	if exist {
		shortUrl := "http://localhost:8080/" + hashValue
		c.JSON(http.StatusOK, apires.Data{
			Base: apires.Base{
				Message: "Url already in database!",
			},
			Data: shortUrl,
		})
		return
	}
	// produce unique hash
	for {
		tempHash := h.CreateSixDigitHash()
		hashExist, err := urlService.CheckHashExist(tempHash)
		if err != nil {
			fmt.Println(err)
		}
		if !hashExist {
			hashValue = tempHash
			break
		}
	}
	// save original url and hash to database => Input: url, hash  Output: error

	hash := edb.Hash{
		Value: hashValue,
	}
	if err := urlService.InsertHash(&hash); err != nil {
		fmt.Println(err)
	}
	url := edb.OriginalUrl{
		HashID: hash.ID,
		Url:    req.OriginalUrl,
	}
	if err := urlService.InsertUrl(&url); err != nil {
		fmt.Println(err)
	}

	// return code, message, data(baseUrl + hash) as response
	shortUrl := "http://localhost:8080/" + hashValue
	c.JSON(http.StatusOK, apires.Data{
		Base: apires.Base{
			Message: "Short URL created!",
		},
		Data: shortUrl,
	})
}
