package routes

import (
	"task/routes/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	user.SetupUserRoutes(app)
}
