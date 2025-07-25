package db

func Tasks(limit int, search string) ([]*Task, error) { // список задач отсортированных по дате.
	var query string
	var args []any

	if search == "" {
		query = `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ?`
		args = []any{limit}
	} else {
		if isDateFormat(search) {
			dateStr, err := convertDateFormat(search)
			if err != nil {
				return nil, err
			}

			query = `SELECT id, date, title, comment, repeat FROM scheduler WHERE date = ? ORDER BY date ASC LIMIT ?`
			args = []any{dateStr, limit}
		} else {
			searchPattern := "%" + search + "%"
			query = `SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date ASC LIMIT ?`
			args = []any{searchPattern, searchPattern, limit}
		}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []*Task{}
	}

	return tasks, nil
}
