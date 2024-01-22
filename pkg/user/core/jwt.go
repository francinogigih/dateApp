package user

type Jwt interface {
	CreateToken(username, email string, id int64) (string, error)
}
