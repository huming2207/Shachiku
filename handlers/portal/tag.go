package portal

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shachiku/common"
	"shachiku/models"
)

func listTags(ctx echo.Context) error {
	tags := &[]models.Tag{}
	db := models.GetDb()

	err := db.Model(tags).Select()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"message": "Failed to load tasks",
		})
	}

	for idx, _ := range *tags {
		err = (*tags)[idx].LoadTasks()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, common.J{
				"message": "Failed to load related tasks",
			})
		}
	}

	return ctx.JSON(http.StatusOK, tags)
}

func addTag(ctx echo.Context) error {
	return nil
}

func listTagDetail(ctx echo.Context) error {
	return nil
}

func deleteTag(ctx echo.Context) error {
	return nil
}
