package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Headliner38/go-project4.git/Desktop/ioBound/pkg/models"
	"github.com/Headliner38/go-project4.git/Desktop/ioBound/pkg/worker"
)

var (
	tasks      = make(map[string]*models.Task) //key - id, value - task
	tasksMutex sync.RWMutex
)

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := models.GenerateTaskID()

	task := models.NewTask(id)
	tasksMutex.Lock()
	tasks[id] = task // добавили задачу в память
	tasksMutex.Unlock()

	go worker.ExecuteTask(task)

	response := map[string]string{"id": id}
	json.NewEncoder(w).Encode(response)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "failed to get task id"})
		return
	}
	tasksMutex.RLock()
	task, exists := tasks[id]
	tasksMutex.RUnlock()

	if !exists {
		writeJson(w, http.StatusNotFound, map[string]string{"error": "task not exists"})
		return
	}
	// ответ для клиента со всеми данными о долгой задаче
	response := map[string]interface{}{
		"id":          task.ID,
		"status":      task.Status,
		"created_at":  task.CreatedAt,
		"started_at":  task.StartedAt,
		"finished_at": task.FinishedAt,
		"duration":    task.TaskDuration().String(),
		"result":      task.Result,
	}

	writeJson(w, http.StatusOK, response)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "failed to get task id"})
		return
	}
	tasksMutex.Lock()
	deleted := models.DeleteTask(tasks, id)
	tasksMutex.Unlock()

	if deleted {
		writeJson(w, http.StatusOK, map[string]string{"OK:": "deleted task with id " + id})
	} else {
		writeJson(w, http.StatusNotFound, map[string]string{"error": "task not found"})
	}
}

func writeJson(w http.ResponseWriter, statusCode int, data any) {
	_ = statusCode
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Printf("ошибка encode: %v\n", err)
	}
}
