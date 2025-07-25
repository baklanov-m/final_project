package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"todo/pkg/db"
	"todo/pkg/nextdate"
)

func writeJSON(w http.ResponseWriter, data any) { // сериализация данных в JSON и отправляет HTTP ответ
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Ошибка сериализации JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func checkDate(task *db.Task) error { // проверка и корректировка даты
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(nextdate.DateFormat)
	}

	t, err := time.Parse(nextdate.DateFormat, task.Date)
	if err != nil {
		return errors.New("дата представлена в формате, отличном от " + nextdate.DateFormat)
	}

	var next string
	if task.Repeat != "" {
		next, err = nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return errors.New("правило повторения указано в неправильном формате: " + err.Error())
		}
	}
	if nextdate.AfterNow(now, t) {
		if task.Repeat == "" {
			task.Date = now.Format(nextdate.DateFormat)
		} else {
			task.Date = next
		}
	}

	return nil
}
