package repositories

import (
	"fmt"

	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

func CreateTaks(inputTask models.Task) {
	fmt.Println(inputTask)
}

func GetTaks() []models.Task {
	db := databases.Conn()
	var tasks []models.Task
	db.Find(&tasks)
	return tasks
}
