package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"shachiku/common"
)

func main() {
	config, err := common.GetConfig().GetSection(common.JwtSection)
	if err != nil {
		log.Panicln(err)
	}

	router := echo.New()

	// Register middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(config.Key(common.JwtSecret).String()),
		SigningMethod: config.Key(common.JwtSignMethod).String(),
		Claims:        common.JwtUserClaims{},
		ContextKey:    "user",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
	}))

	err = router.Start("0.0.0.0:3000")
	if err != nil {
		log.Panicln(err)
	}
}
