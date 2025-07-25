package nextdate // содержит функции для управления повторяющимися задачами.

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) { // определяет следующую дату выполнения повторяющейся задачи
	if repeat == "" {
		return "", errors.New("правило повторения не может быть пустым")
	}

	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", errors.New("некорректная исходная дата: " + err.Error())
	}

	parts := strings.Split(repeat, " ")

	switch parts[0] {
	case "y":
		if len(parts) != 1 {
			return "", errors.New("неверный формат правила для ежегодного повторения")
		}

		for {
			date = date.AddDate(1, 0, 0)
			if AfterNow(date, now) {
				break
			}
		}

	case "d":
		if len(parts) != 2 {
			return "", errors.New("неверный формат правила для повторения по дням")
		}

		interval, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", errors.New("некорректное количество дней: " + err.Error())
		}

		if interval <= 0 || interval > 400 {
			return "", errors.New("количество дней должно быть от 1 до 400")
		}

		for {
			date = date.AddDate(0, 0, interval)
			if AfterNow(date, now) {
				break
			}
		}

	default:
		return "", errors.New("неизвестный тип правила повторения: " + parts[0])
	}

	return date.Format(DateFormat), nil
}

func AfterNow(date, now time.Time) bool {
	dateOnly := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	nowOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return dateOnly.After(nowOnly)
}
