package portal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shachiku/common"
	"shachiku/models"
	"strconv"
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
	tagName := ctx.FormValue("name")
	if tagName == "" {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Tag name cannot be empty",
		})
	}

	db := models.GetDb()
	tag := &models.Tag{Name: tagName}
	_, err := db.Model(tag).Returning("id").Insert()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"message": "Failed to save tag",
		})
	}

	return ctx.JSON(http.StatusOK, tag)
}

func listTagDetail(ctx echo.Context) error {
	tagId, err := strconv.ParseUint(ctx.Param("tagId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Tag ID must be a valid number",
		})
	}

	tag := &models.Tag{ID: uint(tagId)}
	db := models.GetDb()
	err = db.Select(tag)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.J{
			"message": fmt.Sprintf("Cannot find tag with ID %d", tagId),
		})
	}

	return ctx.JSON(http.StatusOK, tag)
}

func deleteTag(ctx echo.Context) error {
	tagId, err := strconv.ParseUint(ctx.Param("tagId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Tag ID must be a valid number",
		})
	}

	tag := &models.Tag{ID: uint(tagId)}
	db := models.GetDb()
	err = db.Delete(tag)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.J{
			"message": fmt.Sprintf("Cannot delete tag with ID %d", tagId),
		})
	}

	return ctx.JSON(http.StatusOK, tag)
}
