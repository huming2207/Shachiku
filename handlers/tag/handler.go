package tag

import "github.com/labstack/echo"

func RegisterTask(rtGroup *echo.Echo) {
	rtGroup.Group("/tag")
	{
		rtGroup.GET("/", listTags)
	}
}

func listTags(ctx echo.Context) error {
	return nil
}
