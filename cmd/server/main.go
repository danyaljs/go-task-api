package main

import (
	"log"
	"net/http"
	"os"
	"task-api/internal/handlers"
	"task-api/internal/middleware"
	"task-api/internal/repository"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	// Initialize repository
	repo := repository.NewInMemoryTaskRepository()

	// Initialize handlers
	taskHandler := handlers.NewTaskHandler(repo)

	// Set up routes with middleware
	http.HandleFunc("/tasks", middleware.Chain(
		taskHandler.HandleTasks,
		middleware.Logger,
		middleware.CORS,
		middleware.Recovery,
		middleware.ContentTypeJSON,
	))

	http.HandleFunc("/tasks/", middleware.Chain(
		taskHandler.HandleTaskByID,
		middleware.Logger,
		middleware.CORS,
		middleware.Recovery,
		middleware.ContentTypeJSON,
	))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if port[0] != ':' {
		port = ":" + port
	}

	log.Printf("🚀 Server starting on http://localhost%s", port)
	log.Printf("📝 Available endpoints:")
	log.Printf("   GET    /tasks       - Get all tasks")
	log.Printf("   POST   /tasks       - Create a new task")
	log.Printf("   GET    /tasks/{id}  - Get task by ID")
	log.Printf("   PUT    /tasks/{id}  - Update task by ID")
	log.Printf("   DELETE /tasks/{id}  - Delete task by ID")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
