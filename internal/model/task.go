package model

import "time"

type Task struct {
	ID        int        `json:"id"`
	Task      string     `json:"task"`
	Message   string     `json:"message"`
	StartAt   time.Time  `json:"start_at"`
	EndAt     *time.Time `json:"end_at"` 
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
