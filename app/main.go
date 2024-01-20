package main

import (
	"context"
	"dateApp/config"
	"fmt"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	ctx := context.Background()
	// define echo backend
	e := echo.New()
	e.Use(mw.Recover())

	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	address := fmt.Sprintf(":%d", config.GetConfig().AppPort)
	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
