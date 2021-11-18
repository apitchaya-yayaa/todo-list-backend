package main

import (
	"todo-list/controller"
	"todo-list/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	storage.NewDB()

	// Routes
	e.GET("/tasks", controller.GetTasks)
	e.GET("/tasks/:id", controller.GetTask)
	e.POST("/tasks", controller.CreateTask)
	e.PUT("/tasks/:id", controller.UpdateTask)
	e.DELETE("/tasks/:id", controller.DeleteTask)

	e.Logger.Fatal(e.Start(":1323"))
}
