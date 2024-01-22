package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Jwt struct {
	secret string
}

func NewJwt(secret string) Jwt {
	return Jwt{secret}
}

type jwtClaims struct {
	UserId   int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"preferred_username"`
	jwt.StandardClaims
}

func (j *Jwt) CreateToken(username, email string, id int64) (string, error) {
	standartClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    username,
	}
	claims := jwtClaims{
		UserId:         id,
		Username:       username,
		Email:          email,
		StandardClaims: standartClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
