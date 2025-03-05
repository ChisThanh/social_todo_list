package common

import "net/http"

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewFullErrorResponse(statusCode int, rootError error, message, log, key string) *AppError {
	return &AppError{StatusCode: statusCode, RootErr: rootError, Message: message, Log: log, Key: key}
}

func NewErrorResponse(err error, msg, log, key string) *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, err, msg, log, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "BD_ERROR")
}

func NewInvalidParamError(err error) *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, err, "invalid parameter", err.Error(), "INVALID_PARAM")
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}
