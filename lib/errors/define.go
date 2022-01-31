package errors

import (
	"net/http"
)

var (
	// ErrInvalidParams when getting invalid params from API request
	ErrInvalidParams = NewErr(http.StatusBadRequest, CODE_INVALID_PARAMS)
	// ErrNoData when data not exists
	ErrNoData = NewErr(http.StatusNotFound, CODE_NOT_EXISTS)
	// ErrDB when getting database error
	ErrDB = NewErr(http.StatusInternalServerError, CODE_DB_ERR)
)
