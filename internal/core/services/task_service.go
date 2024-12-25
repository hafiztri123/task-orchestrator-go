package services

import (
	"context"
	"fmt"

	"github.com/hafiztri123/task-orchestrator-go/internal/core/domain"
	"github.com/hafiztri123/task-orchestrator-go/internal/core/ports"
)

type TaskService struct {
	repository ports.TaskRepository
}

func NewTaskService(repository ports.TaskRepository) *TaskService {
	return &TaskService{
		repository: repository,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, name string, definition map[string]interface{}, maxRetries int) (*domain.Task, error){
	task := domain.NewTask(name, definition, maxRetries)

	if err := s.repository.Save(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	return s.repository.GetByID(ctx, id)
}