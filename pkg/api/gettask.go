package api

import (
	"database/sql"
	"net/http"

	"todo/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) { // обработка GET-запросов для получения задачи по её ID
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeJSON(w, map[string]string{"error": "Задача не найдена"})
		} else {
			writeJSON(w, map[string]string{"error": "Ошибка получения задачи: " + err.Error()})
		}
		return
	}

	writeJSON(w, task)
}
