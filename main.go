package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"shachiku/common"
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

	serverConfig := common.GetConfig().Section(common.ServerSection)
	err := router.Start(serverConfig.Key(common.ServerListen).String())
	if err != nil {
		log.Panicln(err)
	}
}
