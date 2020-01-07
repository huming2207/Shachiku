package user

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"shachiku/common"
)

func RegisterHandler(rtGroup *echo.Group) {
	rtGroup.Group("/user")
	{
		rtGroup.POST("/login", login)
		rtGroup.POST("/register", register)
		rtGroup.POST("/password", changePassword)
	}
}

func login(ctx echo.Context) error {
	return nil
}

func register(ctx echo.Context) error {
	user := &User{}
	err := ctx.Bind(&user)
	if err != nil {
		return err
	}

	passwordStr := ctx.FormValue("password")

	if user.Email == "" || passwordStr == "" || user.Username == "" {
		err = ctx.JSON(http.StatusBadRequest, common.JSON{
			"err": "Empty field detected",
		})

		if err != nil {
			return err
		}

		return nil
	}

	// Generate password hash
	err = user.SetPassword(passwordStr)
	if err != nil {
		err = ctx.JSON(http.StatusInternalServerError, common.JSON{
			"err": fmt.Sprintf("Failed to set password: %v", err),
		})

		if err != nil {
			return err
		}

		return nil
	}

	// Create user
	db := common.GetDb()
	db.Create(&user)

	// Query again to get the ID
	createdUser := &User{}
	db.Where(&User{Username: user.Username, Email: user.Email}).First(&createdUser)

	// Reply with query result
	err = ctx.JSON(http.StatusOK, createdUser)

	if err != nil {
		return err
	}

	return nil
}

func changePassword(ctx echo.Context) error {
	return nil
}
