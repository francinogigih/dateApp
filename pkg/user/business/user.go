package business

import (
	user "dateApp/pkg/user/core"
)

type UserService struct {
	userRepository user.UserRepository
	jwt            user.Jwt
}

func NewUserService(repo user.UserRepository, jwt user.Jwt) UserService {
	return UserService{
		repo,
		jwt,
	}
}

func (s UserService) Register(data *User) (int64, error) {
	userData := user.User{
		Username:   data.Username,
		Password:   data.Password,
		FullName:   data.FullName,
		Email:      data.Email,
		Gender:     data.Gender,
		Phone:      data.Phone,
		BirthPlace: data.BirthPlace,
		BirthDate:  data.BirthDate,
		Location:   data.Location,
	}
	id, err := s.userRepository.Create(&userData)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s UserService) Login(username, password string) (*user.User, string, error) {
	userData, err := s.userRepository.Get(username, password)
	if err != nil {
		return nil, "", err
	}

	token, err := s.jwt.CreateToken(userData.Username, userData.Email, userData.ID)
	if err != nil {
		return nil, "", err
	}

	return userData, token, nil
}
