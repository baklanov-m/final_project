package api

import (
	"net/http"
	"time"

	"todo/pkg/nextdate"
)

func nextDateHandler(w http.ResponseWriter, r *http.Request) { // обрабатывает GET-запросы /api/nextdate
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	nowParam := r.FormValue("now")
	dateParam := r.FormValue("date")
	repeatParam := r.FormValue("repeat")

	if dateParam == "" {
		http.Error(w, "Параметр 'date' обязателен", http.StatusBadRequest)
		return
	}
	if repeatParam == "" {
		http.Error(w, "Параметр 'repeat' обязателен", http.StatusBadRequest)
		return
	}

	var now time.Time
	if nowParam == "" {
		now = time.Now()
	} else {
		var err error
		now, err = time.Parse(nextdate.DateFormat, nowParam)
		if err != nil {
			http.Error(w, "Некорректный формат параметра 'now': "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	result, err := nextdate.NextDate(now, dateParam, repeatParam)
	if err != nil {
		http.Error(w, "Ошибка вычисления следующей даты: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
