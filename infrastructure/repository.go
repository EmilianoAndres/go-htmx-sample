package infrastructure

import (
	"encoding/json"

	domain "github.com/go-htmx-sample/domain"
	"github.com/go-htmx-sample/support/database"
	"github.com/google/uuid"
)

// TaskRepository defines the interface for task data access.
type TaskRepository interface {
	Create(task domain.Task) (domain.Task, error)
	// GetByID(id int) (*domain.Task, error)
	GetAll() ([]domain.Task, error)
	// Update(task *domain.Task) error
	// Delete(id int) error
}

type DefaultTaskRepository struct {
	redisClient database.RedisClient
}

func NewDefaultTaskRepository(redisClient database.RedisClient) DefaultTaskRepository {
	return DefaultTaskRepository{
		redisClient: redisClient,
	}
}

func (r DefaultTaskRepository) Create(task domain.Task) (domain.Task, error) {
	client := r.redisClient.GetClient()
	ctx := r.redisClient.GetContext()

	task.Id = uuid.NewString()

	serialized, _ := json.Marshal(task)

	err := client.Set(ctx, "tasks-"+task.Id, serialized, 0).Err()
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil

}

func (r DefaultTaskRepository) GetAll() ([]domain.Task, error) {

	client := r.redisClient.GetClient()
	ctx := r.redisClient.GetContext()

	var cursor uint64

	keys, cursor, err := client.Scan(ctx, cursor, "tasks-*", 0).Result()

	if err != nil {
		panic(err)
	}

	var returnObject []domain.Task

	for _, key := range keys {
		object, _ := client.Get(ctx, key).Result()

		var deserialized domain.Task

		err = json.Unmarshal([]byte(object), &deserialized)

		returnObject = append(returnObject, deserialized)
	}

	return returnObject, nil
}
