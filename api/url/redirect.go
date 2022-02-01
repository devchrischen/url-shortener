package url

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/entities/edb"
	"github.com/devchrischen/url-shortener/lib/apires"
	"github.com/devchrischen/url-shortener/lib/db"
	"github.com/devchrischen/url-shortener/lib/errors"
	t "github.com/devchrischen/url-shortener/lib/time"
	surl "github.com/devchrischen/url-shortener/services/url"
)

type redirectRequest struct {
	HashValue string `uri:"hashValue" binding:"required"`
}

func Redirect(c *gin.Context) {
	// validate request
	var req redirectRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errors.Throw(c, errors.ErrInvalidParams.SetError(err))
		return
	}

	// check param is valid hash
	match, _ := regexp.MatchString("[A-Za-z0-9]{6}", req.HashValue)
	if !match {
		errors.Throw(c, errors.ErrInvalidParams)
		return
	}
	// query db to check if the hash exist
	urlService := surl.New(db.DB)
	hash := edb.Hash{}
	if err := urlService.GetHash(&hash, req.HashValue); err != nil {
		errors.Throw(c, err)
		return
	}
	// check if the hash is not expired
	expired := t.CheckHashExpired(hash.CreatedAt)
	if expired {
		// delete url record
		if err := urlService.DeleteUrl(hash.ID); err != nil {
			errors.Throw(c, err)
			return
		}
		// delete hash record
		if err := urlService.DeleteHash(hash.ID); err != nil {
			errors.Throw(c, err)
			return
		}
		// throw error
		errors.Throw(c, errors.ErrExpired)
		return
	}

	// query db to get original url
	url := edb.OriginalUrl{}
	if err := urlService.GetUrl(&url, hash.ID); err != nil {
		errors.Throw(c, err)
		return
	}

	// return code, message, data(redirect url) as response
	originalUrl := url.Url
	c.JSON(http.StatusOK, apires.Data{
		Base: apires.Base{
			Code:    errors.CODE_OK,
			Message: "Find redirect URL successfully!",
		},
		Data: originalUrl,
	})
}
