package model

import (
	"database/sql"
)

type Task struct {
	ID        int64
	Task      string
	Message   sql.NullString
	StartAt   string
	EndAt     sql.NullString
	CreatedAt string
	UpdatedAt string
}
