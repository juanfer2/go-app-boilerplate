package models

import (
	"gorm.io/gorm"
)

// Types
type Task struct {
	gorm.Model
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []Task

// Persistence
var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}
