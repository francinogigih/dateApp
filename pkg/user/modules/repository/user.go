package repository

import (
	// "os/user"
	"errors"
	"time"

	appErr "dateApp/pkg/common/http"
	user "dateApp/pkg/user/core"

	"github.com/jackc/pgx/v5/pgconn"
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

func (repo *PostgresDBRepository) toCore(data User) *user.User {
	return &user.User{
		ID:         data.ID,
		Username:   data.Username,
		Password:   data.Password,
		FullName:   data.FullName,
		Email:      data.Email,
		Phone:      data.Phone,
		Gender:     data.Gender,
		BirthPlace: data.BirthPlace,
		BirthDate:  data.BirthDate,
		Location:   data.Location,
	}
}

func (repo *PostgresDBRepository) Create(data *user.User) (int64, error) {
	userData := repo.newUserData(data)
	result := repo.db.Create(&userData)

	if result.Error != nil {
		if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
			if pgError.Code == "23505" {
				return 0, errors.New(appErr.ErrAlreadyRegistered)
			}
			return 0, errors.New(appErr.ErrInternalServer)
		}
		return 0, errors.New(appErr.ErrInternalServer)
	}
	return userData.ID, nil
}

func (repo *PostgresDBRepository) Get(username, password string) (*user.User, error) {
	var userData User
	result := repo.db.Where("username = ?", username).Where("password = ?", password).First(&userData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New(appErr.ErrInvalidUser)
		}
		return nil, errors.New(appErr.ErrInternalServer)
	}

	return repo.toCore(userData), nil
}
