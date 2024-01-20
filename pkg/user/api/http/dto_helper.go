package http

import (
	"dateApp/pkg/user/api/http/request"
	"dateApp/pkg/user/business"
	"time"
)

func newUserData(params *request.UserRegistration) *business.User {
	birthDate, _ := time.Parse(time.RFC3339, params.BirthDate)
	return &business.User{
		Username:   params.Username,
		Password:   params.Password,
		FullName:   params.FullName,
		Email:      params.Email,
		Gender:     params.Gender,
		Phone:      params.Phone,
		BirthPlace: params.BirthPlace,
		BirthDate:  birthDate,
		Location:   params.Location,
	}
}
