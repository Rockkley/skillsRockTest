package handlers

import (
	"github.com/gofiber/fiber/v2"
	"skillsRockTest/internal/dto"
	"skillsRockTest/internal/models"
	"skillsRockTest/internal/storage"
)

type TaskHandler struct {
	repo storage.TaskRepositoryInterface
}

func NewTaskHandler(repo storage.TaskRepositoryInterface) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	ctx := c.Context()

	var taskDTO dto.TaskCreateDTO
	if err := c.BodyParser(&taskDTO); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid data"})
	}

	if err := taskDTO.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation error", "details": err.Error()})
	}

	task := models.Task{
		Title:       taskDTO.Title,
		Description: taskDTO.Description,
	}

	if err := h.repo.CreateTask(ctx, task); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error creating task"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "the task successfully created"})
}

func (h *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	ctx := c.Context()
	tasks, err := h.repo.GetAllTasks(ctx)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error getting tasks"})
	}

	return c.JSON(tasks)
}
func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	ctx := c.Context()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	var taskDTO dto.TaskUpdateDTO
	if err := c.BodyParser(&taskDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid data"})
	}

	if err = taskDTO.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation error", "details": err.Error()})

	}

	task := models.Task{
		Title:       taskDTO.Title,
		Description: taskDTO.Description,
		Status:      string(taskDTO.Status),
	}

	if err := h.repo.UpdateTask(ctx, id, task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error updating the task"})
	}

	return c.JSON(fiber.Map{"message": "the task updated successfully"})
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	ctx := c.Context()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := h.repo.DeleteTask(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error deleting the task"})
	}

	return c.JSON(fiber.Map{"message": "the task has been deleted"})
}
