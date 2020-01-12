package portal

import (
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	group.GET("", getUser)
	group.POST("/password", changePassword)

	// Tag group
	tagGroup := group.Group("/tag")
	tagGroup.GET("", listTags)
	tagGroup.GET("/:tagId", listTagDetail)
	tagGroup.POST("", addTag)
	tagGroup.DELETE("/:tagId", deleteTag)

	// Task group
	taskGroup := group.Group("/task")
	taskGroup.POST("", addTask)
	taskGroup.GET("", getAllTasks)
	taskGroup.GET("/:taskId", getOneTask)
	taskGroup.DELETE("/:taskId", removeTask)
}

func changePassword(ctx echo.Context) error {
	token := ctx.Get(common.JwtSection).(*jwt.Token)
	claims := token.Claims.(*models.JwtUserClaims)

	db := models.GetDb()
	user := &models.User{}
	db.First(&user, claims.UserID)

	password := ctx.FormValue("password")
	err := validation.Validate(password, validation.Required, validation.Length(6, 64))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Password length must be between 6 to 64 characters",
		})
	}

	err = user.SetPassword(password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"message": "Failed to generate password hash",
		})
	}
	db.Save(&user)

	return ctx.JSON(http.StatusOK, common.J{
		"message": "Password updated",
	})
}

func getUser(ctx echo.Context) error {
	token := ctx.Get(common.JwtSection).(*jwt.Token)
	claims := token.Claims.(*models.JwtUserClaims)

	db := models.GetDb()
	user := &models.User{}
	db.First(&user, claims.UserID)

	err := user.LoadRelatedTasks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"message": "Failed to load related tasks",
		})
	}

	return ctx.JSON(http.StatusOK, user)
}
