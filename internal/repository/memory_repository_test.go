package repository

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/robertov8/task_list/internal/models"
)

func TestInMemoryTaskRepository_GetAll(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task1 := createTestTask("Tarefa 1", "Descrição 1", false)
	task2 := createTestTask("Tarefa 2", "Descrição 2", true)

	repo.Add(task1)
	repo.Add(task2)

	tasks := repo.GetAll()

	if len(tasks) != 2 {
		t.Errorf("Esperado 2 tarefas, obtido %d", len(tasks))
	}
}

func TestInMemoryTaskRepository_GetByID(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task := createTestTask("Tarefa Test", "Descrição Test", false)
	repo.Add(task)

	foundTask, found := repo.GetByID(task.ID)
	if !found {
		t.Errorf("Tarefa com ID %s não encontrada", task.ID)
	}

	if foundTask.Title != task.Title {
		t.Errorf("Esperado título '%s', obtido '%s'", task.Title, task.ID)
	}

	_, found = repo.GetByID("id-inexistente")
	if found {
		t.Error("Não deveria encontrar tarefa com ID inexistente")
	}
}

func TestInMemoryTaskRepository_GetByStatus(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task1 := createTestTask("Tarefa 1", "Descrição 1", false)
	task2 := createTestTask("Tarefa 2", "Descrição 2", true)
	task3 := createTestTask("Tarefa 3", "Descrição 3", true)

	repo.Add(task1)
	repo.Add(task2)
	repo.Add(task3)

	completedTasks := repo.GetByStatus(true)
	pendingTasks := repo.GetByStatus(false)

	if len(completedTasks) != 2 {
		t.Errorf("Esperando 2 tarefas completas, obtido %d", len(completedTasks))
	}

	if len(pendingTasks) != 1 {
		t.Errorf("Esperando 1 tarefas pendente, obtido %d", len(pendingTasks))
	}
}

func TestInMemoryTaskRepository_Add(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task := createTestTask("Nova tarefa", "Descrição da nova tarefa", false)

	addedTask := repo.Add(task)
	if addedTask.ID != task.ID {
		t.Errorf("IDs não correspodem após adição,esperado %s obtido %s", task.ID, addedTask.ID)
	}

	allTasks := repo.GetAll()
	if len(allTasks) != 1 {
		t.Errorf("Esperado 1 tarefa após adição, obtido %d", len(allTasks))
	}
}

func TestInMemoryTaskRepository_Update(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task := createTestTask("Tarefa Original", "Descrição Original", false)
	repo.Add(task)

	updatedTask := models.Task{
		Title:       "Tarefa Atualizada",
		Description: "Descrição Atualizada",
		Done:        true,
	}

	resultTask, found := repo.Update(task.ID, updatedTask)

	if !found {
		t.Errorf("Tarefa com ID %s não encontrada para atualização", task.ID)
	}

	if resultTask.Title != updatedTask.Title {
		t.Errorf("Título não atualizado: esperado '%s', obtido '%s'", resultTask.Title, updatedTask.Title)
	}

	if resultTask.Description != updatedTask.Description {
		t.Errorf("Descrição não atualizado: esperado '%s', obtido '%s'", resultTask.Description, updatedTask.Description)
	}

	if resultTask.Done != updatedTask.Done {
		t.Errorf("Status não atualizado: esperado '%v', obtido '%v'", resultTask.Done, updatedTask.Done)
	}

	_, found = repo.Update("id-inexistente", updatedTask)
	if found {
		t.Error("Não deveria encontrar tarefas com ID inexistente para atualização")
	}
}

func TestInMemoryTaskRepository_Delete(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task := createTestTask("Tarefa para Excluir", "Descrição da tarefa para elixir", false)
	repo.Add(task)

	if len(repo.GetAll()) != 1 {
		t.Fatal("Tarefa não foi adicionada corretamente para o teste de exclusão")
	}

	deleted := repo.Delete(task.ID)

	if !deleted {
		t.Error("Operação de exclusão retornou falso para ID existente")
	}

	if len(repo.GetAll()) != 0 {
		t.Error("Tarefa não foi excluída corretamente")
	}

	deleted = repo.Delete("id-inexistente")
	if deleted {
		t.Error("Operação de exclusão retornou verdadeiro para ID inexistente")
	}
}

func createTestTask(title, description string, done bool) models.Task {
	return models.Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Done:        done,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
