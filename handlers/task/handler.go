package task

import "github.com/labstack/echo"

func RegisterTask(rtGroup *echo.Group) {
	rtGroup.Group("/task")
	{
		rtGroup.POST("/", addTask)
		rtGroup.GET("/", getTaskList)
		rtGroup.GET("/:taskId", getTaskDetail)
		rtGroup.DELETE("/:taskId", removeTask)
	}
}

func addTask(ctx echo.Context) error {
	return nil
}

func getTaskList(ctx echo.Context) error {
	return nil
}

func getTaskDetail(ctx echo.Context) error {
	return nil
}

func removeTask(ctx echo.Context) error {
	return nil
}
