package storage

import (
	"context"
	"skillsRockTest/internal/models"
)

type TaskRepositoryInterface interface {
	CreateTask(ctx context.Context, task models.Task) error
	GetAllTasks(ctx context.Context) ([]models.Task, error)
	UpdateTask(ctx context.Context, id int, task models.Task) error
	DeleteTask(ctx context.Context, id int) error
}
