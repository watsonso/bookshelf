package controllers

import (
	"github.com/watsonso/bookshelf/models"
)

type Task struct {
}

func NewTask() Task {
	return Task{}
}

func (c Task) Get(n int) interface{} {
	repo := models.NewTaskRepository()
	task := repo.GetByID(n)
	return task
}

func (c Task) GetAll() interface{} {
	repo := models.NewTaskRepository()
	tasks := repo.GetAll()
	return tasks
}

func (c Task) Create(text string) {
	repo := models.NewTaskRepository()
	repo.Create(text)
}
