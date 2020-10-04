package repositories

import (
	"log"

	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

func CreateTaks(inputTask models.Task) models.Task {
	db := databases.Conn()
	task := inputTask
	db.Create(&task)
	log.Println(task)
	return task
}

func GetTaks() []models.Task {
	db := databases.Conn()
	var tasks []models.Task
	db.Find(&tasks)
	return tasks
}
