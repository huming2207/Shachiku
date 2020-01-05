package task

import "github.com/gin-gonic/gin"

func RegisterTask(rtGroup *gin.RouterGroup) {
	rtGroup.Group("/task")
	{
		rtGroup.POST("/", addTask)
		rtGroup.GET("/", getTaskList)
		rtGroup.GET("/:taskId", getTaskDetail)
		rtGroup.DELETE("/:taskId", removeTask)
	}
}

func addTask(ctx *gin.Context) {

}

func getTaskList(ctx *gin.Context) {

}

func getTaskDetail(ctx *gin.Context) {

}

func removeTask(ctx *gin.Context) {

}
