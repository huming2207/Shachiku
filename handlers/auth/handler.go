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

type authResponse struct {
	Message   string       `json:"message"`
	Token     string       `json:"token,omitempty"`
	ExpiresAt int64        `json:"expires_at,omitempty"`
	User      *models.User `json:"user,omitempty"`
}

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

func generateJwt(userName string, userId uint) (string, int64, error) {
	expiresAt := time.Now().Add(time.Hour).Unix()
	claims := &models.JwtUserClaims{
		UserName: userName,
		UserID:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte(jwtConfig.Key(common.JwtSecret).String()))
	if err != nil {
		return "", -1, err
	}

	return tokenStr, expiresAt, nil
}

func login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	// Check empty or not
	if username == "" || password == "" {
		return ctx.JSON(http.StatusBadRequest, &authResponse{
			Message: "Empty or invalid request",
		})
	}

	// Find auth by auth name or email
	user := &models.User{}
	db := models.GetDb()
	err := db.Model(user).Where("username = ?", username).WhereOr("email = ?", username).Select()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &authResponse{
			Message: "User not found",
		})
	}

	// Validate password
	match, err := user.CheckPassword(password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &authResponse{
			Message: "Password validation failed",
		})
	}

	if !match {
		return ctx.JSON(http.StatusUnauthorized, &authResponse{
			Message: "Password incorrect",
		})
	}

	// Generate JWT
	tokenStr, expiresAt, err := generateJwt(username, user.ID)
	return ctx.JSON(http.StatusOK, &authResponse{
		Message:   "OK",
		Token:     tokenStr,
		ExpiresAt: expiresAt,
	})
}

func register(ctx echo.Context) error {
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	passwordStr := ctx.FormValue("password")

	// Validate auth name
	err := validation.Validate(username, validation.Required, validation.Length(3, 50))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &authResponse{
			Message: "Username is invalid",
		})
	}

	// Validate password
	err = validation.Validate(passwordStr, validation.Required, validation.Length(6, 64))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &authResponse{
			Message: "Password must be between 6 to 64 characters",
		})
	}

	// Validate email
	err = validation.Validate(email, validation.Required, is.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &authResponse{
			Message: "Email address is invalid",
		})
	}

	// Generate password hash
	user := &models.User{Username: username, Email: email}
	err = user.SetPassword(passwordStr)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &authResponse{
			Message: "Failed to set password",
		})
	}

	// Create auth
	db := models.GetDb()
	_, err = db.Model(user).Returning("id").Insert()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &authResponse{
			Message: "Failed to create user. Username/Email may be duplicated.",
		})
	}

	// Reply with query result
	tokenStr, expiresAt, err := generateJwt(username, user.ID)
	return ctx.JSON(http.StatusOK, &authResponse{
		Message:   "OK",
		Token:     tokenStr,
		ExpiresAt: expiresAt,
		User:      user,
	})
}
