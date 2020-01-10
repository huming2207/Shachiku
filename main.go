package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"shachiku/handlers/auth"
	"shachiku/handlers/portal"
)

func main() {
	router := echo.New()

	// Register middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	auth.RegisterHandler(router)
	portal.RegisterHandler(router)

	err := router.Start("0.0.0.0:3000")
	if err != nil {
		log.Panicln(err)
	}
}
