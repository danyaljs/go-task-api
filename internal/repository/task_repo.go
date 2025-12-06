package repository

import (
	"sync"
	"task-api/internal/models"
	"time"
)

type TaskRepository interface {
	Create(title, description string) (*models.Task, error)
	GetAll() ([]*models.Task, error)
	GetByID(id int) (*models.Task, error)
	Update(id int, updates *models.UpdateTaskRequest) (*models.Task, error)
	Delete(id int) error
}

type InMemoryTaskRepository struct {
	mu     sync.RWMutex
	tasks  map[int]*models.Task
	nextID int
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[int]*models.Task),
		nextID: 1,
	}
}

func (r *InMemoryTaskRepository) Create(title, description string) (*models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	task := &models.Task{
		ID:          r.nextID,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	r.tasks[r.nextID] = task
	r.nextID++

	return task, nil
}

func (r *InMemoryTaskRepository) GetAll() ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *InMemoryTaskRepository) GetByID(id int) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]

	if !exists {
		return nil, models.ErrTaskNotFound
	}

	return task, nil
}

func (r *InMemoryTaskRepository) Update(id int, updates *models.UpdateTaskRequest) (*models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]

	if !exists {
		return nil, models.ErrTaskNotFound
	}

	if updates.Title != nil {
		task.Title = *updates.Title
	}
	if updates.Description != nil {
		task.Description = *updates.Description
	}
	if updates.Completed != nil {
		task.Completed = *updates.Completed
	}

	task.UpdatedAt = time.Now()

	return task, nil

}

func (r *InMemoryTaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return models.ErrTaskNotFound
	}

	delete(r.tasks, id)

	return nil
}
