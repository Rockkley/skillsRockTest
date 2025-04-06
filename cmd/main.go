package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"skillsRockTest/config"
	"skillsRockTest/internal/storage/repository"
	"skillsRockTest/internal/transport/handlers"
	"skillsRockTest/internal/transport/routes"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.GetConfig()
	connStr := cfg.GetDBConnStr()

	db, err := repository.NewDatabase(connStr)
	if err != nil {

		log.Fatal(err)
	}
	defer db.Conn.Close(context.Background())

	err = repository.RunMigrations(db.Conn)
	if err != nil {

		log.Fatal(err)
	}

	taskRepo := repository.NewTaskRepository(db.Conn)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	app := fiber.New(
		fiber.Config{AppName: cfg.AppName},
	)
	routes.SetupRoutes(app, taskHandler)

	port := cfg.AppPort
	log.Printf("server is running at port:%s", port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("gracefully shutting down...")
		if err := app.Shutdown(); err != nil {
			log.Fatalf("error on gracefull shutdown: %v", err)
		}
	}()

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("starting server failed: %v", err)
	}
}
