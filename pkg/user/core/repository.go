package user

type UserRepository interface {
	Create(userData *User) (int64, error)
	Get(username, password string) (*User, error)
}
