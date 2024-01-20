package business

import user "dateApp/pkg/user/core"

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService(repo user.UserRepository) UserService {
	return UserService{
		repo,
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
