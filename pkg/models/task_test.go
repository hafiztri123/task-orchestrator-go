package models

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	name := "Test Task"
	desc := "Test Description"
	priority := PriorityHigh
	task := NewTask(name, desc, priority)

	if task.Name != name {
		t.Errorf("Expected task name %s, got %s", name, task.Name)
	}

	if task.Status != TaskStatusPending {
		t.Errorf("Expected status PENDING, got %s", task.Status)
	}

	if task.Priority != priority {
		t.Errorf("Expected priority %d, got %d", priority, task.Priority)
	}
}

func TestTaskExpiry(t *testing.T){
	task := NewTask("Test", "Description", PriorityMedium)

	if task.HasExpired(){
		t.Error("Task without deadline should not be expired")
	}

	future := time.Now().Add(time.Hour)
	task.Deadline = &future
	if task.HasExpired(){
		t.Error("Task with future deadline should not be expired")
	}

	past := time.Now().Add(-time.Hour)
	task.Deadline = &past
	if !task.HasExpired(){
		t.Error(("Task with past deadline should be expired"))
	}
}