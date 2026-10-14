package main

import (
	"log"
	"os"

	"go-backend-demo/internal/handler"
	"go-backend-demo/internal/repository"
	"go-backend-demo/internal/router"
	"go-backend-demo/internal/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Dependency injection thủ công: repo -> service -> handler -> router.
	// Muốn đổi repo in-memory sang Postgres chỉ cần thay dòng dưới đây.
	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	r := router.New(taskHandler)

	log.Printf("server listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
