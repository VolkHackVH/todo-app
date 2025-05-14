package handlers

import "github.com/VolkHackVH/todo-list/internal/db"

type Handlers struct {
	User *UserHandler
	Task *TasksHandler
}

func NewHandler(db *db.Queries) *Handlers {
	return &Handlers{
		User: newUserHandler(db),
		Task: newTaskHandler(db),
	}
}
