package http

import (
	common "dateApp/pkg/common/http"
	"net/http"

	"dateApp/pkg/user/api/http/request"
	"dateApp/pkg/user/api/http/response"
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
	userId, err := h.service.Register(dto)
	if err != nil {
		errResp := common.RenderErrorResponse(err)
		return c.JSON(errResp.Code, errResp)
	}

	resp := common.NewCreatedSuccessResponse(userId)
	return c.JSON(resp.Code, resp)
}

func (h *Handler) Login(c echo.Context) error {
	req := new(request.Login)
	c.Bind(req)

	userData, token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		errResp := common.RenderErrorResponse(err)
		return c.JSON(errResp.Code, errResp)
	}

	resp := response.NewLoginResponse(userData, token)
	return c.JSON(http.StatusOK, resp)
}
