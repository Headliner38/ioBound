package api

import "net/http"

func Init() {
	http.HandleFunc("/api/tasks/create", createTaskHandler)
	http.HandleFunc("/api/tasks/get", getTaskHandler)
	http.HandleFunc("/api/tasks/delete", deleteTaskHandler)
}
