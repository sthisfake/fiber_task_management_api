package task

import (
	"task/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(app *fiber.App) {
	task := app.Group("/tasks")
	task.Post("/", handlers.CreateTask)
	task.Put("/:id", handlers.UpdateTask)
	task.Delete("/:id", handlers.DeleteTask)
	task.Get("/:id", handlers.GetTask)
	task.Get("/:userId", handlers.GetTaskList)
}
