package tag

import "github.com/gin-gonic/gin"

func RegisterTask(rtGroup *gin.RouterGroup) {
	rtGroup.Group("/tag")
	{
		rtGroup.GET("/", listTags)
	}
}

func listTags(ctx *gin.Context) {

}
