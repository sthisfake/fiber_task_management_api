package user

import (
	"task/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/create_user", handlers.CreateUser)
}
