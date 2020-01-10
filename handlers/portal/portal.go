package portal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
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
		SigningMethod: "HS512",
		Claims:        &models.JwtUserClaims{},
		ContextKey:    common.JwtSection,
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
	}))

	group.POST("/password", changePassword)

	// Tag group
	tagGroup := group.Group("/tag")
	tagGroup.GET("/", listTags)

	// Task group
	taskGroup := group.Group("/task")
	taskGroup.POST("/", addTask)
	taskGroup.GET("/", getTaskList)
	taskGroup.GET("/:taskId", getTaskDetail)
	taskGroup.DELETE("/:taskId", removeTask)
}

func changePassword(ctx echo.Context) error {
	token := ctx.Get(common.JwtSection).(*jwt.Token)
	claims := token.Claims.(*models.JwtUserClaims)

	db := models.GetDb()
	user := &models.User{}
	db.First(&user, claims.UserID)

	err := user.SetPassword(ctx.FormValue("password"))
	if err != nil {
		return err
	}
	db.Save(&user)

	return ctx.JSON(http.StatusOK, common.JSON{
		"message": "Password updated",
	})
}
