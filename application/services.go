package application

import (
	"github.com/go-htmx-sample/domain"
	infrastructure "github.com/go-htmx-sample/infrastructure"
)

// TaskService defines the interface for task-related business logic.
type TaskService interface {
	CreateTask(name string) error
	// GetTaskByID(id int) (*domain.Task, error)
	GetAllTasks() ([]domain.Task, error)
	// UpdateTask(task *domain.Task) error
	// DeleteTask(id int) error
}

type DefaultTaskService struct {
	taskRepo infrastructure.TaskRepository
}

// NewTaskService creates a new task service.
func NewTaskService(taskRepo infrastructure.TaskRepository) DefaultTaskService {
	return DefaultTaskService{
		taskRepo: taskRepo,
	}
}

func (s DefaultTaskService) CreateTask(name string) error {
	print(name)
	return nil
}

func (s DefaultTaskService) GetAllTasks() ([]domain.Task, error) {
	tasks, _ := s.taskRepo.GetAll()

	return tasks, nil
}

// Implement TaskService methods here.
// Example: CreateTask, GetTaskByID, GetAllTasks, UpdateTask, DeleteTask
