package user

import "time"

type User struct {
	ID         int64
	Username   string
	Password   string
	FullName   string
	Email      string
	Gender     string
	Phone      string
	BirthPlace string
	BirthDate  time.Time
	Location   string
}
