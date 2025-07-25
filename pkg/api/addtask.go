package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todo/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) { // отвечает за обработку POST-запросов для создания новых задач
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка десериализации JSON: " + err.Error()})
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

	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка добавления задачи: " + err.Error()})
		return
	}

	writeJSON(w, map[string]string{"id": strconv.FormatInt(id, 10)})
}
