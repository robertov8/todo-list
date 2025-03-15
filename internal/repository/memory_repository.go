package repository

import (
	"sync"
	"time"

	"github.com/robertov8/task_list/internal/models"
)

type InMemoryTaskRepository struct {
	tasks map[string]models.Task
	mutex sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]models.Task),
	}
}

func (r *InMemoryTaskRepository) GetAll() []models.Task {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	taskList := make([]models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		taskList = append(taskList, task)
	}

	return taskList
}

func (r *InMemoryTaskRepository) GetByID(id string) (models.Task, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	task, found := r.tasks[id]

	return task, found
}

func (r *InMemoryTaskRepository) GetByStatus(done bool) []models.Task {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var taskList []models.Task
	for _, task := range r.tasks {
		if task.Done == done {
			taskList = append(taskList, task)
		}
	}

	return taskList
}

func (r *InMemoryTaskRepository) Add(task models.Task) models.Task {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.tasks[task.ID] = task
	return task
}

func (r *InMemoryTaskRepository) Update(id string, task models.Task) (models.Task, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	existingTask, found := r.tasks[id]
	if !found {
		return models.Task{}, false
	}

	if task.Title != "" {
		existingTask.Title = task.Title
	}

	if task.Description != "" {
		existingTask.Description = task.Description
	}

	existingTask.Done = task.Done
	existingTask.UpdatedAt = time.Now()

	r.tasks[id] = existingTask

	return existingTask, true
}

func (r *InMemoryTaskRepository) Delete(id string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, found := r.tasks[id]
	if !found {
		return false
	}

	delete(r.tasks, id)
	return true
}
