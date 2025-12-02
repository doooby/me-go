package repository

import (
	"database/sql"
	"me-go/db"
	app "me-go/internal"
	"me-go/internal/model"
	"time"
)

func CreateTask(taskName string, message sql.NullString, startAt time.Time) (int64, error) {
	query := "INSERT INTO tasks (task, message, start_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	nowStr := app.TimeToStr(time.Now())
	result, err := db.DB.Exec(query, taskName, message, app.TimeToStr(startAt), nowStr, nowStr)
	if err == nil {
		id, err := result.LastInsertId()
		if err == nil && id > 0 {
			return id, nil
		}
	}
	return 0, err
}

func FindUnfinishedId() (int64, error) {
	query := "SELECT id FROM tasks WHERE id = (SELECT MAX(id) FROM tasks) AND end_at IS NULL;"
	var id int64
	err := db.DB.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FindTaskById(id int64) (model.Task, error) {
	query := "SELECT * FROM tasks WHERE id = ?"
	var task model.Task
	err := db.DB.QueryRow(query, id).Scan(
		&task.ID,
		&task.Task,
		&task.Message,
		&task.StartAt,
		&task.EndAt,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func UpdateTaskEndAt(id int64, end_at time.Time) error {
	query := "UPDATE tasks SET end_at = ?, updated_at = ? WHERE id = ?"
	result, err := db.DB.Exec(query, app.TimeToStr(end_at), app.TimeToStr(time.Now()), id)
	if err == nil {
		rowsAffected, err := result.RowsAffected()
		if err == nil && rowsAffected == 1 {
			return nil
		}
	}
	return err
}

// func ListTasks(showAll bool) ([]model.Task, error) {
// 	var query string
// 	if showAll {
// 		query = `SELECT id, caption, text, start_time, end_time FROM tasks`
// 	} else {
// 		query = `SELECT id, caption, text, start_time, end_time FROM tasks WHERE end_time IS NULL`
// 	}
//
// 	rows, err := db.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
//
// 	var tasks []model.Task
// 	for rows.Next() {
// 		var t model.Task
// 		// We scan EndTime into a sql.NullTime or pointer to handle NULLs
// 		// Here we strictly check for the pointer scan
// 		err := rows.Scan(&t.ID, &t.Task, &t.Message, &t.StartAt, &t.EndAt, &t.CreatedAt, &t.UpdatedAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tasks = append(tasks, t)
// 	}
// 	return tasks, nil
// }
//
// func CompleteTask(id int) error {
// 	query := `UPDATE tasks SET end_time = ? WHERE id = ?`
// 	_, err := db.DB.Exec(query, time.Now(), id)
// 	return err
// }
