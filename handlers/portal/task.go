package portal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"shachiku/common"
	"shachiku/models"
	"strconv"
	"time"
)

func addTask(ctx echo.Context) error {
	title := ctx.FormValue("title")
	location := ctx.FormValue("location")
	comment := ctx.FormValue("comment")

	// Validate title
	if title == "" {
		return ctx.JSON(http.StatusBadRequest, common.JSON{
			"message": "Title can't be empty",
		})
	}

	// Retrieve start_at and end_at times
	startAtTs, err := strconv.ParseInt(ctx.FormValue("start_at"), 10, 64)
	if err != nil {
		return err
	}

	endAtTs, err := strconv.ParseInt(ctx.FormValue("end_at"), 10, 64)
	if err != nil {
		return err
	}

	startAt := time.Unix(startAtTs, 0)
	endAt := time.Unix(endAtTs, 0)

	// Get current user
	jwtToken := ctx.Get(common.JwtSection).(*jwt.Token)
	jwtClaims := jwtToken.Claims.(*models.JwtUserClaims)

	// Save to database
	db := models.GetDb()
	task := &models.Task{
		Title:    title,
		Location: location,
		Comment:  comment,
		StartAt:  startAt,
		EndAt:    endAt,
		OwnerID:  jwtClaims.UserID,
	}

	err = db.Save(&task).Error
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, common.JSON{
		"message": "ok",
		"task":    task,
	})
}

func getAllTasks(ctx echo.Context) error {
	tasks := &[]models.Task{}
	db := models.GetDb()
	err := db.Find(&tasks).Error
	if err != nil {
		return err
	}

	for idx, _ := range *tasks {
		err = (*tasks)[idx].LoadOwner()
		if err != nil {
			return err
		}
	}

	return ctx.JSON(http.StatusOK, &tasks)
}

func getOneTask(ctx echo.Context) error {
	db := models.GetDb()
	taskId := ctx.Param("taskId")
	task := &models.Task{}

	err := db.First(&task, taskId).Error
	if err != nil {
		return err
	}

	err = task.LoadOwner()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &task)
}

func removeTask(ctx echo.Context) error {
	db := models.GetDb()
	taskId := ctx.Param("taskId")
	task := &models.Task{}

	err := db.Delete(&task, taskId).Error
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &task)
}
