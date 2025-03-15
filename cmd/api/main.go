package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/robertov8/task_list/internal/handlers"
	"github.com/robertov8/task_list/internal/middleware"
	"github.com/robertov8/task_list/internal/models"
	"github.com/robertov8/task_list/internal/repository"
)

func main() {
	taskRepo := repository.NewInMemoryTaskRepository()

	task1 := models.Task{
		ID:          uuid.New().String(),
		Title:       "Aprender Golang",
		Description: "Estudar a linguagem Go e criar uma aplicação exemplo",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task2 := models.Task{
		ID:          uuid.New().String(),
		Title:       "Comprar mantimentos",
		Description: "Ir ao supermercado comprar leite, pão e ovos",
		Done:        true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	taskRepo.Add(task1)
	taskRepo.Add(task2)

	taskHandler := handlers.NewTaskHandler(taskRepo)

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/api/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", taskHandler.GetTask).Methods("GET")
	r.HandleFunc("/api/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	port := 4000
	fmt.Printf("Servidor iniciado na porta http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
