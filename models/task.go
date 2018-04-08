package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//var engine *xorm.Engine

func init() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:@/gin-todo")
	if err != nil {
		panic(err)
	}
}

type Task struct {
	ID   int    `json:"id" xorm:"'id'"`
	Text string `json:"text" xorm:"'text'"`
}

type Tasks []Task

type TaskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return TaskRepository{}
}

func (m TaskRepository) Create(text string) {
	var task = Task{Text: text}
	engine.Insert(&task)
}

func (m TaskRepository) GetByID(id int) *Task {
	var task = Task{ID: id}
	has, _ := engine.Get(&task)
	if has {
		return &task
	}
	return nil
}

func (m TaskRepository) GetAll() Tasks {
	var task = Task{}
	var tasks = Tasks{}
	rows, err := engine.Rows(task)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&task)
		tasks = append(tasks, task)
	}
	return tasks
}
