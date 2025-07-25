package db

import "fmt"

type Task struct {
	ID      string `json:"id"`      // id задачи
	Date    string `json:"date"`    // дата выполнения
	Title   string `json:"title"`   // задача
	Comment string `json:"comment"` // коммент к задаче
	Repeat  string `json:"repeat"`  // правила повторения
}

func AddTask(task *Task) (int64, error) { // добавление задачи
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetTask(id string) (*Task, error) { // получение задачи по id
	task := &Task{}

	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`

	err := db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func UpdateTask(task *Task) error { // обновленеие существующей задачи
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("задача с ID %s не найдена", task.ID)
	}

	return nil
}

func DeleteTask(id string) error { // удаляет задачу по id
	query := `DELETE FROM scheduler WHERE id = ?`

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("задача с ID %s не найдена", id)
	}

	return nil
}

func UpdateDate(nextDate string, id string) error { // обновление даты задачи
	query := `UPDATE scheduler SET date = ? WHERE id = ?`

	res, err := db.Exec(query, nextDate, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("задача с ID %s не найдена", id)
	}

	return nil
}
