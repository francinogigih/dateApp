package business

import "time"

type User struct {
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
