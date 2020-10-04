package services

import (
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/repositories"
)

func CreateTakService(inputTaks models.Task) {
	repositories.CreateTaks(inputTaks)
}

func GetTasksService() []models.Task {
	return repositories.GetTaks()
}
