package routes

import (
	"skillsRockTest/internal/transport/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, taskHandler *handlers.TaskHandler) {
	api := app.Group("/api")

	api.Post("/tasks", taskHandler.CreateTask)
	api.Get("/tasks", taskHandler.GetAllTasks)
	api.Put("/tasks/:id", taskHandler.UpdateTask)
	api.Delete("/tasks/:id", taskHandler.DeleteTask)

}
