package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"skillsRockTest/internal/models"
)

const (
	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
)

type TaskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task models.Task) error {
	err := r.db.QueryRow(ctx,
		"INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id",
		task.Title, task.Description, StatusNew).Scan(&task.ID)
	if err != nil {
		log.Printf("error while creating task: %v", err)
		return err
	}
	return nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, description, status FROM tasks")
	if err != nil {
		log.Printf("error while getting all tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			log.Printf("getAllTasks error scanning tasks: %v", err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, id int, task models.Task) error {
	query := `UPDATE tasks SET
	title = CASE WHEN $1 = '' THEN title ELSE $1 END,
	description = CASE WHEN $2 = '' THEN description ELSE $2 END,
	status = CASE WHEN $3 = '' THEN status ELSE $3 END,
	updated_at = now()
    WHERE id = $4;`

	_, err := r.db.Exec(ctx, query, task.Title, task.Description, task.Status, id)
	if err != nil {
		log.Printf("error updating task: %v", err)
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Printf("error deleting task: %v", err)
		return err
	}
	return nil
}
