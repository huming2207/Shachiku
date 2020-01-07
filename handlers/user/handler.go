package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shachiku/common"
)

func RegisterHandler(rtGroup *gin.RouterGroup) {
	rtGroup.Group("/user")
	{
		rtGroup.POST("/login", login)
		rtGroup.POST("/register", register)
		rtGroup.POST("/password", changePassword)
	}
}

func login(ctx *gin.Context) {

}

func register(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	if userName == "" || password == "" || email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "Empty field detected",
		})
	}

	user := &User{
		Username: userName,
		Email:    email,
	}

	// Generate password hash
	err := user.SetPassword(password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": fmt.Sprintf("Failed to set password: %v", err),
		})
	}

	// Create user
	db := common.GetDb()
	db.Create(&user)

	// Query again to get the ID
	createdUser := &User{}
	db.Where(&User{Username: userName, Email: email}).First(&createdUser)

	// Reply with query result
	ctx.JSON(http.StatusOK, gin.H{
		"id":        createdUser.ID,
		"user_name": userName,
		"email":     email,
	})
}

func changePassword(ctx *gin.Context) {

}
