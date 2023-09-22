package infrastructure

import (
	domain "github.com/go-htmx-sample/domain"
)

// TaskRepository defines the interface for task data access.
type TaskRepository interface {
	Create(task domain.Task) error
	// GetByID(id int) (*domain.Task, error)
	GetAll() ([]domain.Task, error)
	// Update(task *domain.Task) error
	// Delete(id int) error
}

type DefaultTaskRepository struct{}

func NewDefaultTaskRepository() DefaultTaskRepository {
	return DefaultTaskRepository{}
}

func (r DefaultTaskRepository) Create(task domain.Task) error {
	return nil
}

func (r DefaultTaskRepository) GetAll() ([]domain.Task, error) {
	return []domain.Task{{
		ID:   1,
		Name: "Task 1",
	}, {
		ID:   2,
		Name: "Task 2",
	}}, nil
}
