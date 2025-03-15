package repository

import "github.com/robertov8/task_list/internal/models"

type TaskRepository interface {
	GetAll() []models.Task
	GetByID(id string) (models.Task, bool)
	GetByStatus(done bool) []models.Task
	Add(task models.Task) models.Task
	Update(id string, task models.Task) (models.Task, bool)
	Delete(id string) bool
}
