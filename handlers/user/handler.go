package user

import "github.com/gin-gonic/gin"

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

}

func changePassword(ctx *gin.Context) {

}
