package url

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/api/url/request"
	"github.com/devchrischen/url-shortener/entities/edb"
	"github.com/devchrischen/url-shortener/lib/apires"
	"github.com/devchrischen/url-shortener/lib/db"
	"github.com/devchrischen/url-shortener/lib/errors"
	surl "github.com/devchrischen/url-shortener/services/url"
)

func Redirect(c *gin.Context) {
	// validate request
	r := request.RedirectRequest{}
	req, err := r.Validate(c)
	if err != nil {
		errors.Throw(c, err)
	}

	// query db to check if the hash exist
	urlService := surl.New(db.DB)
	hash := edb.Hash{}
	if err := urlService.GetHash(&hash, req.HashValue); err != nil {
		errors.Throw(c, err)
		return
	}
	// check if the hash is not expired
	expired := urlService.CheckHashExpired(hash.CreatedAt)
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
