package repository

import (
	// "os/user"
	"errors"
	"time"

	appErr "dateApp/pkg/common/http"
	user "dateApp/pkg/user/core"

	"gorm.io/gorm"
)

type PostgresDBRepository struct {
	db *gorm.DB
}

func NewPostgresDBRepository(db *gorm.DB) *PostgresDBRepository {
	return &PostgresDBRepository{
		db,
	}
}

type User struct {
	ID         int64  `gorm:"primary_key,autoIncrement"`
	Username   string `gorm:"uniqueIndex:username"`
	Password   string
	FullName   string
	Email      string `gorm:"uniqueIndex:email"`
	Gender     string
	Phone      string
	BirthPlace string
	BirthDate  time.Time
	Location   string
}

func (repo *PostgresDBRepository) newUserData(userData *user.User) User {
	return User{
		Username:   userData.Username,
		Password:   userData.Password,
		FullName:   userData.FullName,
		Email:      userData.Email,
		Phone:      userData.Phone,
		Gender:     userData.Gender,
		BirthPlace: userData.BirthPlace,
		BirthDate:  userData.BirthDate,
		Location:   userData.Location,
	}
}

func (repo *PostgresDBRepository) Create(data *user.User) (int64, error) {
	userData := repo.newUserData(data)
	result := repo.db.Save(&userData)

	if result.Error != nil {
		return 0, errors.New(appErr.ErrInternalServer)
	}
	return userData.ID, nil
}
