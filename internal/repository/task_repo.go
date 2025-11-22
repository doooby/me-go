package repository

import (
	"me-go/db"
	"me-go/internal/model"
	"time"
)

func CreateTask(caption, text string) error {
	query := `INSERT INTO tasks (caption, text, start_time) VALUES (?, ?, ?)`
	_, err := db.DB.Exec(query, caption, text, time.Now())
	return err
}

func ListTasks(showAll bool) ([]model.Task, error) {
	var query string
	if showAll {
		query = `SELECT id, caption, text, start_time, end_time FROM tasks`
	} else {
		query = `SELECT id, caption, text, start_time, end_time FROM tasks WHERE end_time IS NULL`
	}

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		// We scan EndTime into a sql.NullTime or pointer to handle NULLs
		// Here we strictly check for the pointer scan
		err := rows.Scan(&t.ID, &t.Task, &t.Message, &t.StartAt, &t.EndAt, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func CompleteTask(id int) error {
	query := `UPDATE tasks SET end_time = ? WHERE id = ?`
	_, err := db.DB.Exec(query, time.Now(), id)
	return err
}
