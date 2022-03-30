package domain

import (
	"encoding/json"
	"net/http"
)

type CustomError struct {
	HTTPCode     int    `json:"-"`
	ErrorCode    int    `json:"errorCode"`
	MessageError string `json:"messageError"`
}

func (ce CustomError) Error() string {
	return ce.MessageError
}

var (
	UnknownErr = CustomError{
		HTTPCode:     http.StatusInternalServerError,
		ErrorCode:    1,
		MessageError: "something unexpected happened",
	}
)

func (ce CustomError) String() string {
	data, err := json.Marshal(ce)
	if err != nil {
		return `{"message":"error unknown","errorCode":0}`
	}

	return string(data)
}
