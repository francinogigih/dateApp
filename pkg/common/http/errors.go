package common

import (
	"net/http"
)

const (
	ErrNotFound       = "err_not_found"
	ErrInternalServer = "err_internal_server"
)

// DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message"`
	// Internal error  `json:"-"`
}

// CreatedResponse default payload response
type CreatedSuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload int64  `json:"payload"`
}

// ErrorResponse error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message"`
}

func NewInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		500,
		InternalErrStatus,
		"Internal server error",
	}
}

func RenderErrorResponse(err error) (resp ErrorResponse) {
	switch err.Error() {
	case ErrNotFound:
		resp = ErrorResponse{Status: NotFoundStatus, Code: http.StatusNotFound, Message: "Not Found"}
	default:
		resp = ErrorResponse{Status: InternalErrStatus, Code: http.StatusInternalServerError, Message: "Internal server error"}
	}
	return resp
}
