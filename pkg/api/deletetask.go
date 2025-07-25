package api

import (
	"net/http"

	"todo/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) { // обработка DELETE- запросов для удаления задач
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	err := db.DeleteTask(id) // удаляем
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, map[string]any{})
}
