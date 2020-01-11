package auth

import (
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
	"github.com/labstack/echo/v4"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"shachiku/common"
	"shachiku/models"
	"time"
)

var jwtConfig *ini.Section

func RegisterHandler(router *echo.Echo) {
	var err error
	jwtConfig, err = common.GetConfig().GetSection(common.JwtSection)
	if err != nil {
		log.Fatalln(err)
	}

	group := router.Group("/auth")

	group.POST("/register", register)
	group.POST("/login", login)
}

func login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	// Check empty or not
	if username == "" || password == "" {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Empty or invalid request",
		})
	}

	// Find auth by auth name or email
	user := &models.User{}
	db := models.GetDb()
	db.First(&user, "username = ? OR email = ?", username, username)

	// Validate password
	match, err := user.CheckPassword(password)
	if err != nil {
		return err
	}

	if !match {
		return ctx.JSON(http.StatusUnauthorized, common.J{
			"message": "Password incorrect",
		})
	}

	// Setup claims
	expiresAt := time.Now().Add(time.Hour).Unix()
	claims := &models.JwtUserClaims{
		UserName: username,
		UserID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	var tokenStr string
	tokenStr, err = token.SignedString([]byte(jwtConfig.Key(common.JwtSecret).String()))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, common.J{
		"message": "OK",
		"token":   tokenStr,
		"expires": expiresAt,
	})
}

func register(ctx echo.Context) error {
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	passwordStr := ctx.FormValue("password")

	// Validate auth name
	err := validation.Validate(username, validation.Required, validation.Length(3, 50))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"error": "Username is invalid",
		})
	}

	// Validate password
	err = validation.Validate(passwordStr, validation.Required, validation.Length(6, 64))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"error": "Password must be between 6 to 64 characters",
		})
	}

	// Validate email
	err = validation.Validate(email, validation.Required, is.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"error": "Email address is invalid",
		})
	}

	// Generate password hash
	user := &models.User{Username: username, Email: email}
	err = user.SetPassword(passwordStr)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"error": "Failed to set password",
		})
	}

	// Create auth
	db := models.GetDb()
	err = db.Create(&user).Error
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"error": "Failed to create user. Username/Email may be duplicated.",
		})
	}

	// Query again to get the ID
	createdUser := &models.User{}
	db.Where(&models.User{Username: user.Username, Email: user.Email}).First(&createdUser)

	// Reply with query result
	return ctx.JSON(http.StatusOK, createdUser)
}
