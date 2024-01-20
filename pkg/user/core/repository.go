package user

type UserRepository interface {
	Create(userData *User) (int64, error)
}
