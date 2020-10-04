package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/services"
)

// Create Task
func CreateTaks(c *fiber.Ctx) error {
	return c.JSON("CreateTask")
}

func GetTaks(c *fiber.Ctx) error {
	Tasks := services.GetTasksService()
	return c.JSON(Tasks)

}

func UpdateTaks(c *fiber.Ctx) error {
	return c.JSON("UpdateTask")

}

func DeleteTaks(c *fiber.Ctx) error {
	return c.JSON("DeleteTask")

}
