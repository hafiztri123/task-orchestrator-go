package ports

import (
	"context"

	"github.com/hafiztri123/task-orchestrator-go/internal/core/domain"
)

type TaskRepository interface {
	Save(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id string) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	List(ctx context.Context, filter map[string]interface{}) ([]*domain.Task, error)
}