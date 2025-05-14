package service

import (
	"context"
	"errors"

	"github.com/VolkHackVH/todo-list/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskService struct {
	Db *db.Queries
}

func NewTaskService(db *db.Queries) *TaskService {
	return &TaskService{
		Db: db,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, text string, created_at pgtype.Timestamptz) (db.Task, error) {
	if len(text) < 1 || len(text) > 500 {
		return db.Task{}, errors.New("text is not within the allowed characters")
	}

	task, err := s.Db.CreateTask(ctx, db.CreateTaskParams{
		Text:      text,
		CreatedAt: created_at,
	})
	if err != nil {
		return db.Task{}, err
	}

	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.Db.DeleteTask(ctx, int32(id))
}

func (s *TaskService) GetTaskByID(ctx context.Context, id int) (db.Task, error) {
	return s.Db.GetTaskByID(ctx, int32(id))
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]db.Task, error) {
	tasks, err := s.Db.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, text string) error {
	if len(text) <= 8 {
		return errors.New("text cannot be less than 9 characters long")
	}

	return s.Db.UpdateTask(ctx, db.UpdateTaskParams{
		ID:   int32(id),
		Text: text,
	})
}
