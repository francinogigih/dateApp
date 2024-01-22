package response

import (
	user "dateApp/pkg/user/core"
)

type Login struct {
	User        user.User `json:"user"`
	AccessToken string    `json:"accessToken"`
}

// type LoginPayload struct {
// 	user.User
// 	Token
// }

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload Login  `json:"payload"`
}

func NewLoginResponse(result *user.User, token string) *LoginResponse {
	var ResultResponse LoginResponse
	loginResp := Login{
		User:        *result,
		AccessToken: token,
	}

	ResultResponse.Code = 200
	ResultResponse.Message = "Success"

	ResultResponse.Payload = loginResp

	return &ResultResponse
}
