package api

import (
	"net/http"

	"todo/pkg/db"
)

type TasksResp struct { // структура ответа для API endpoint
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) { // обработка GET запроса к /api/tasks.
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	search := r.FormValue("search")

	tasks, err := db.Tasks(50, search)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка получения задач: " + err.Error()})
		return
	}

	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
