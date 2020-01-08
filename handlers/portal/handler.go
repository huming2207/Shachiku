package portal

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/ini.v1"
	"log"
	"shachiku/common"
	"shachiku/models"
)

var jwtConfig *ini.Section

func RegisterHandler(router *echo.Echo) {
	var err error
	jwtConfig, err = common.GetConfig().GetSection(common.JwtSection)
	if err != nil {
		log.Fatalln(err)
	}

	group := router.Group("/portal")

	group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(jwtConfig.Key(common.JwtSecret).String()),
		SigningMethod: jwtConfig.Key(common.JwtSignMethod).String(),
		Claims:        models.JwtUserClaims{},
		ContextKey:    "user",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
	}))

	group.POST("/password", changePassword)
}

func changePassword(ctx echo.Context) error {
	return nil
}
