package models

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	StatusPending   = "pending"
	StatusRunning   = "running"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
)

type Task struct {
	ID         string     `json:"id"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	Result     string     `json:"result"`
	mutex      sync.RWMutex
}

func NewTask(id string) *Task {
	return &Task{
		ID:        id,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
}
func DeleteTask(tasks map[string]*Task, id string) bool {
	if _, exists := tasks[id]; exists {
		delete(tasks, id)
		return true
	}
	return false
}

func GenerateTaskID() string {
	return uuid.New().String()
}

func (t *Task) UpdateTaskStatus(newStatus string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.Status == newStatus { // если статус не изменился, то ничего не делаем
		return
	}
	t.Status = newStatus
	now := time.Now()
	switch newStatus {
	case StatusRunning:
		t.StartedAt = &now
	case StatusCompleted:
		t.FinishedAt = &now
	case StatusFailed:
		t.FinishedAt = &now
	}
}

func (t *Task) TaskDuration() time.Duration {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	if t.StartedAt == nil {
		return 0
	}
	if t.FinishedAt != nil { // если уже завершилась, то возвращаем разницу между завершением и началом
		return t.FinishedAt.Sub(*t.StartedAt)
	}
	return time.Since(*t.StartedAt) // если не завершилась, то возвращаем время с начала
}
