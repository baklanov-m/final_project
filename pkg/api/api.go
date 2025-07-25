package api

import "net/http"

func Init() { // выполняет регистрацию всех обработчиков API.
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", doneTaskHandler)
}
