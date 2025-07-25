package db

import (
	"errors"
	"strings"
	"time"
)

func isDateFormat(s string) bool { // проверяет, совпадает ли строка с форматом даты DD.MM.YYYY.
	if len(s) != 10 {
		return false
	}

	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return false
	}

	if len(parts[0]) != 2 || len(parts[1]) != 2 || len(parts[2]) != 4 {
		return false
	}

	for _, part := range parts {
		for _, char := range part {
			if char < '0' || char > '9' {
				return false
			}
		}
	}

	return true
}

func convertDateFormat(dateStr string) (string, error) { // преобразует дату из формата DD.MM.YYYY в формат YYYYMMDD.
	t, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		return "", errors.New("некорректный формат даты")
	}

	return t.Format("20060102"), nil
}
