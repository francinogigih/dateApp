package http

import "github.com/labstack/echo/v4"

// RegisterPath Register V1 API path
func RegisterPath(e *echo.Echo, h *Handler) {
	if h == nil {
		panic("item controller cannot be nil")
	}
	e.POST("v1/register", h.RegisterUser)
	e.POST("v1/login", h.Login)
}
