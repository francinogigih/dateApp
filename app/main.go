package main

import (
	"context"
	"dateApp/config"
	"fmt"

	pg "dateApp/pkg/common/modules/db"
	userAPIHttp "dateApp/pkg/user/api/http"
	userService "dateApp/pkg/user/business"
	userRepository "dateApp/pkg/user/modules/repository"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func newUserService(db *gorm.DB) userService.UserService {
	userRepository := userRepository.NewPostgresDBRepository(db)
	userService := userService.NewUserService(userRepository)
	return userService
}

func main() {
	db := pg.DatabaseConnection()

	userService := newUserService(db)
	userHandler := userAPIHttp.NewHandler(userService)
	ctx := context.Background()
	// define echo backend
	e := echo.New()
	e.Use(mw.Recover())

	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	userAPIHttp.RegisterPath(e, userHandler)

	address := fmt.Sprintf(":%d", config.GetConfig().AppPort)
	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
