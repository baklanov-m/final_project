package api

import (
	"encoding/json"
	"net/http"

	"todo/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) { //  отвечает за обработку PUT-запросов для обновления задач.
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка десериализации JSON: " + err.Error()})
		return
	}

	if task.ID == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор задачи"})
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, map[string]any{})
}
