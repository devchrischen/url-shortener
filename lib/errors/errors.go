package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/devchrischen/url-shortener/lib/apires"
)

type (
	ICustomError interface {
		GetHttpCode() int
		GetCode() int
		GetMsgCode() int
		GetVerbs() map[string]string
		GetData() interface{}
		HasErrors() bool
		GetErrors() map[string][]interface{}
		Reset()
		error
	}

	CustomMsg string
)

const (
	CODE_UNKNOWN_ERR       = -1
	CODE_OK                = 0
	CODE_INVALID_PARAMS    = 10
	CODE_DB_ERR            = 15
	CODE_NOT_EXISTS        = 16
	CODE_DUPLICATE_KEY     = 17
	CODE_SHORT_URL_EXPIRED = 21
)

var ErrCodeMsgMap = map[int]string{
	CODE_INVALID_PARAMS:    "Invalid params from request",
	CODE_SHORT_URL_EXPIRED: "Short url is expired",
}

func GetMessage(code int) string {
	if message, ok := ErrCodeMsgMap[code]; ok {
		return message
	}
	return ""
}

func Throw(c *gin.Context, err error) {

	if gorm.IsRecordNotFoundError(err) {
		DBError(c, err)
		return
	}

	switch e := err.(type) {
	case *mysql.MySQLError, gorm.Errors:
		DBError(c, e)
	case ICustomError:
		CustomError(c, e)
	default:
		Error(c, http.StatusInternalServerError, CODE_UNKNOWN_ERR, e)
	}
}

func CustomError(c *gin.Context, err ICustomError) {

	httpCode := err.GetHttpCode()
	code := err.GetCode()
	errMsg := err.Error()

	if err.HasErrors() {
		Error(c, httpCode, code, err.GetErrors())
		err.Reset()
		return
	}

	if _, ok := ErrCodeMsgMap[code]; ok {
		Error(c, httpCode, code, CustomMsg(errMsg))
		err.Reset()
		return
	}

	Error(c, http.StatusInternalServerError, CODE_UNKNOWN_ERR, CustomMsg(errMsg))
	err.Reset()
}

func DBError(c *gin.Context, err error) {

	if gorm.IsRecordNotFoundError(err) {
		Error(c, http.StatusNotFound, CODE_NOT_EXISTS, "Data not found")
		return
	}

	myerr, ok := err.(*mysql.MySQLError)
	if !ok {
		Error(c, http.StatusInternalServerError, CODE_DB_ERR, err)
		return
	}

	if myerr.Number == 1062 {
		Error(c, http.StatusBadRequest, CODE_DUPLICATE_KEY, "Duplicate data already exists")
	} else if myerr.Number == 1452 {
		Error(c, http.StatusBadRequest, CODE_INVALID_PARAMS, "Reference ID not exists")
	} else {
		Error(c, http.StatusInternalServerError, CODE_DB_ERR, err)
	}
}

func Error(c *gin.Context, httpCode int, code int, err interface{}) {
	var msg string
	switch err := err.(type) {
	case CustomMsg:
		// Error message already handled by CustomError
		msg = string(err)
	case error:
		msg = getMsgByCode(c, code, err.Error())
	case string:
		msg = getMsgByCode(c, code, err)
	}

	c.JSON(httpCode, apires.Base{
		Code:    code,
		Message: msg,
	})

	c.Abort()
}

func getMsgByCode(c *gin.Context, code int, msg string) string {

	switch code {

	case CODE_NOT_EXISTS:
		msg = "Data not found"

	case CODE_INVALID_PARAMS:
		if msg == "" {
			msg = "Invalid input parameter"
		}
	}
	return msg
}
