package domain

import (
	"time"
	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskStatusPending TaskStatus = "PENDING"
	TaskStatusRunning TaskStatus = "RUNNING"
	TaskStatusCompleted TaskStatus = "COMPLETED"
	TaskStatusFailed TaskStatus = "FAILED"
)

type Task struct {
	ID uuid.UUID
	Name string
	Status TaskStatus
	Definition map[string]interface{}
	RetryCount int
	MaxRetries int
	CreatedAt time.Time
	StartedAt *time.Time
	CompletedAt *time.Time
	Error string
}

func NewTask(name string, definition map[string]interface{}, maxRetries int) *Task {
	return &Task{
		ID: uuid.New(),
		Name: name,
		Status: TaskStatusPending,
		Definition: definition,
		MaxRetries: maxRetries,
		CreatedAt: time.Now(),
	}
}