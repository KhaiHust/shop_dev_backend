package common

import "net/http"

type ResponseCommon struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ErrorResponse struct {
	HTTPCode    int
	ServiceCode int
	Message     string
}

const (
	CannotBindJson      = 40000001
	CannotCreateNewUser = 40000002
	InvalidLogin        = 400000003
)

var errResponseMap = map[int]ErrorResponse{
	CannotBindJson: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: CannotBindJson,
		Message:     "Can not bind JSON",
	},
	CannotCreateNewUser: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: CannotCreateNewUser,
		Message:     "Can not create new user",
	},
}

func GetErrorResponse(code int) ErrorResponse {
	if val, ok := errResponseMap[code]; ok {
		return val
	}

	return ErrorResponse{
		HTTPCode:    http.StatusInternalServerError,
		ServiceCode: code,
		Message:     http.StatusText(http.StatusInternalServerError),
	}
}
