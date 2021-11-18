package controller

import (
	"net/http"
	"todo-list/model"
	"todo-list/storage"

	"github.com/labstack/echo/v4"
)

// get all tasks
func GetTasks(c echo.Context) error {
	db := storage.GetDBInstance()
	tasks := []model.Task{}

	if err := db.Find(&tasks).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}

// get task by id
func GetTask(c echo.Context) error {
	db := storage.GetDBInstance()
	id := c.Param("id")
	task := model.Task{}

	if err := db.First(&task, id).Error; err != nil {
		return c.String(http.StatusNotFound, "Task not found!")
	}

	return c.JSON(http.StatusOK, task)
}

// Create task
func CreateTask(c echo.Context) error {
	db := storage.GetDBInstance()

	task := model.Task{}

	if err := c.Bind(&task); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := db.Create(&task).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

// Update task
func UpdateTask(c echo.Context) error {
	db := storage.GetDBInstance()
	id := c.Param("id")
	task := model.Task{}
	form := model.Task{}

	if err := db.Find(&task, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := db.Model(&task).Select("*").Omit("id").Updates(&form).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, task)
}

// Delete task
func DeleteTask(c echo.Context) error {
	db := storage.GetDBInstance()
	id := c.Param("id")
	task := model.Task{}

	if err := db.Find(&task, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := db.Delete(&task).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
