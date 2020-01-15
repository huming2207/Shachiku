package portal

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"shachiku/common"
	"shachiku/models"
	"strconv"
	"strings"
	"time"
)

func addTask(ctx echo.Context) error {
	title := ctx.FormValue("title")
	location := ctx.FormValue("location")
	comment := ctx.FormValue("comment")
	tags := strings.Split(ctx.FormValue("tags"), ",")
	var tagIds []uint64

	// Parse Tag ID
	if len(tags) != 0 {
		for _, tagIdStr := range tags {
			tagId, err := strconv.ParseUint(tagIdStr, 10, 64)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, common.J{
					"message": fmt.Sprintf("Cannot parse Tag ID %s", tagIdStr),
				})
			} else {
				tagIds = append(tagIds, tagId)
			}
		}
	}

	// Validate title
	if title == "" {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Title can't be empty",
		})
	}

	// Retrieve start_at and end_at times
	startAtTs, err := strconv.ParseInt(ctx.FormValue("start_at"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Failed to parse start time",
		})
	}

	endAtTs, err := strconv.ParseInt(ctx.FormValue("end_at"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Failed to parse end time",
		})
	}

	startAt := time.Unix(startAtTs, 0)
	endAt := time.Unix(endAtTs, 0)

	// Get current user
	jwtToken := ctx.Get(common.JwtSection).(*jwt.Token)
	jwtClaims := jwtToken.Claims.(*models.JwtUserClaims)

	// Save to database
	db := models.GetDb()
	user := &models.User{ID: jwtClaims.UserID}
	err = db.Select(user)

	task := &models.Task{
		Title:    title,
		Location: location,
		Comment:  comment,
		StartAt:  startAt,
		EndAt:    endAt,
	}

	_, err = db.Model(task).Returning("id").Insert()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Failed to save task to database",
		})
	}

	// Create the owner
	task.People = []*models.Role{
		{UserID: jwtClaims.UserID, TaskID: task.ID, Level: models.Owner},
	}

	err = db.Insert(task.People[0])
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": "Failed to save task to database",
		})
	}

	// Add tags
	for _, tagId := range tagIds {
		err = db.Insert(&models.TagTask{TaskID: task.ID, TagID: uint(tagId)})
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, common.J{
				"message": fmt.Sprintf("Cannot bind tag with ID %d", tagId),
			})
		}
	}

	err = task.LoadPeople()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"message": "Failed to assign owner",
		})
	}

	err = task.LoadTags()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"message": "Failed to assign tags",
		})
	}

	return ctx.JSON(http.StatusOK, task)
}

func getAllTasks(ctx echo.Context) error {
	tasks := &[]models.Task{}
	db := models.GetDb()
	err := db.Model(tasks).Select()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"message": "Failed to load tasks",
		})
	}

	for idx, _ := range *tasks {
		err = (*tasks)[idx].LoadPeople()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, common.J{
				"message": "Failed to load owner",
			})
		}
	}

	return ctx.JSON(http.StatusOK, tasks)
}

func getOneTask(ctx echo.Context) error {
	db := models.GetDb()
	taskId, err := strconv.ParseUint(ctx.Param("taskId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": fmt.Sprintf("Task ID parameter is not a valid number"),
		})
	}

	task := &models.Task{ID: uint(taskId)}
	err = db.Select(task)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.J{
			"message": fmt.Sprintf("Task ID %s was not found", taskId),
		})
	}

	err = task.LoadPeople()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.J{
			"message": "Failed to load owner",
		})
	}

	return ctx.JSON(http.StatusOK, task)
}

func removeTask(ctx echo.Context) error {
	db := models.GetDb()
	taskId, err := strconv.ParseUint(ctx.Param("taskId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.J{
			"message": fmt.Sprintf("Task ID parameter is not a valid number"),
		})
	}

	task := &models.Task{ID: uint(taskId)}
	err = db.Delete(task)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.J{
			"message": fmt.Sprintf("Failed to delete task %d", taskId),
		})
	}

	return ctx.JSON(http.StatusOK, task)
}
