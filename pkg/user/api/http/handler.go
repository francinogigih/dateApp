package http

import (
	common "dateApp/pkg/common/http"

	"dateApp/pkg/user/api/http/request"
	"dateApp/pkg/user/business"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service business.UserService
}

// NewHandler Construct user API handler
func NewHandler(service business.UserService) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) RegisterUser(c echo.Context) error {
	req := new(request.UserRegistration)
	c.Bind(req)

	dto := newUserData(req)
	code, err := h.service.Register(dto)
	if err != nil {
		errResp := common.RenderErrorResponse(err)
		return c.JSON(errResp.Code, errResp)
	}

	resp := common.NewCreatedSuccessResponse(code)
	return c.JSON(resp.Code, resp)
}
