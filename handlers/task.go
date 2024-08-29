package handlers

import (
	"task/database"
	"task/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {

	// check the validation of request body
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	// check if the user id is valid
	var user models.User
	if err := database.DB.First(&user, "id = ?", task.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user is not exist"})
	}

	// create the task in the db
	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong , cant create task"})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	var task models.Task

	// checking if a task exist with this id
	if err := database.DB.First(&task, "id = ?", taskID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no task found with this id"})
	}

	// check the validation of the request body
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	// save the changes to the db
	if err := database.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update task"})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	// checking if a task exist with this id
	taskID := c.Params("id")
	if err := database.DB.Delete(&models.Task{}, "id = ?", taskID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task did not found"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetTask(c *fiber.Ctx) error {
	// checking if a task exist with this id
	taskID := c.Params("id")
	var task models.Task
	if err := database.DB.First(&task, "id = ?", taskID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func GetTaskList(c *fiber.Ctx) error {
	// parsing request body
	var user models.GetUserTasks
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	// geting list of the task by user
	var tasks []models.Task
	if err := database.DB.Where("user_id = ?", user.UserId).Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not get tasks"})
	}

	return c.Status(fiber.StatusOK).JSON(tasks)
}
