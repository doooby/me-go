package repository

import (
	"database/sql"
	"fmt"
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
	query := "SELECT id FROM tasks WHERE id = (" +
		"SELECT id FROM tasks ORDER BY id DESC LIMIT 1" +
		") AND end_at IS NULL;"
	var id int64
	err := db.DB.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FindTaskById(id int64) (model.Task, error) {
	query := "SELECT * FROM tasks WHERE id = ?;"
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

func GetLastTask() (model.Task, error) {
	query := "SELECT * FROM tasks ORDER BY id DESC LIMIT 1;"
	var task model.Task
	err := db.DB.QueryRow(query).Scan(
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
	query := "UPDATE tasks SET end_at = ?, updated_at = ? WHERE id = ?;"
	result, err := db.DB.Exec(query, app.TimeToStr(end_at), app.TimeToStr(time.Now()), id)
	if err == nil {
		rowsAffected, err := result.RowsAffected()
		if err == nil && rowsAffected == 1 {
			return nil
		}
	}
	return err
}

func ListTasks(pagination db.Pagination) ([]model.Task, error) {
	query := fmt.Sprintf("SELECT * FROM tasks ORDER BY id DESC %s;", pagination.SqlFragment())
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]model.Task, 0, pagination.PerPage)
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Task,
			&task.Message,
			&task.StartAt,
			&task.EndAt,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
