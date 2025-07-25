package api

import (
	"database/sql"
	"net/http"
	"time"

	"todo/pkg/db"
	"todo/pkg/nextdate"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) { // обработка POST-запросов для пометки задач как выполненных
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка удаления задачи: " + err.Error()})
			return
		}
	} else {
		now := time.Now()
		nextDate, err := nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка вычисления следующей даты: " + err.Error()})
			return
		}
		err = db.UpdateDate(nextDate, id)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка обновления даты: " + err.Error()})
			return
		}
	}

	writeJSON(w, map[string]any{})
}
