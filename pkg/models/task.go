package models

import (
	"time"
)

type TaskStatus string

const (
	TaskStatusPending TaskStatus = "PENDING"
	TaskStatusRunning TaskStatus = "RUNNING"
	TaskStatusCompleted TaskStatus = "COMPLETED"
	TaskStatusFailed TaskStatus = "FAILED"
)

type Priority int

const (
	PriorityLow Priority = 1
	PriorityMedium Priority = 2
	PriorityHigh Priority = 3
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Priority Priority `json:"priority"`
	Status TaskStatus `json:"status"`
	Payload []byte `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	StartedAt *time.Time `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Deadline *time.Time `json:"deadline"`
	RetryCount int `json:"retry_count"`
	MaxRetries int `json:"max_retries"`
	Error string `json:"error"`
}

func NewTask(name string, description string, priority Priority) *Task {
	now := time.Now()
	return &Task{
		ID: generateID(),//TODO: Implement GenerateID
		Name: name,
		Description: description,
		Priority: priority,
		Status: TaskStatusPending,
		CreatedAt: now,
		UpdatedAt: now,
		RetryCount: 0,
		MaxRetries: 3,
	}	
}

func (t *Task) HasExpired() bool {
	return t.Deadline != nil && time.Now().After(*t.Deadline)
}

func (t *Task) canRetry() bool {
	return t.RetryCount < t.MaxRetries
}

