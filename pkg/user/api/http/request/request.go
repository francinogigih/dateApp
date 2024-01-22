package request

type UserRegistration struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	FullName   string `json:"fullName"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	BirthPlace string `json:"birthPlace"`
	BirthDate  string `json:"birthDate"`
	Location   string `json:"location"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
