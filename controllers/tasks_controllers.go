package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/services"
	"github.com/juanfer2/api-rest-go/validations"
)

func getDataTaks(c *fiber.Ctx) models.Task {
	newTaks := new(models.Task)
	log.Println(newTaks)
	body := c.BodyParser(newTaks)
	log.Println(body)
	return *newTaks
}

// Create Task
func CreateTasks(c *fiber.Ctx) error {
	var p models.Task
	p = getDataTaks(c)
	err := validations.ValidateCreateTask(p)
	if err != nil {
		return c.Status(422).JSON(err)
	}
	data := services.CreateTaskService(p)
	return c.JSON(data)
}

func GetTasks(c *fiber.Ctx) error {
	tasks := services.GetTasksService()
	return c.JSON(tasks)

}

func UpdateTasks(c *fiber.Ctx) error {
	return c.JSON("UpdateTask")

}

func DeleteTasks(c *fiber.Ctx) error {
	return c.JSON("DeleteTask")

}
