package handlers

import (
	"task/database"
	"task/models"
	"task/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	// read request body and parse it to see if its valid
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cant hash the password",
		})
	}
	user.Password = hashedPassword

	// create a user in the db
	if err := database.DB.Create(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong, can't create user",
		})
	}

	// return the created user
	return c.Status(fiber.StatusCreated).JSON(user)
}
