package model

import (
	"database/sql"
	"fmt"
	app "me-go/internal"
	"time"
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

const dateTimeLayout = "06-01-02 15:04"
const timeLayout = "15:04"

func (task Task) DurationText() (string, error) {
	startAt, err := app.StrToTime(task.StartAt)
	if err != nil {
		return "", fmt.Errorf("invalid start_at [%d] '%s': %v", task.ID, task.StartAt, err)
	}
	startAt = startAt.Local()

	if task.EndAt.Valid {
		endAt, err := app.StrToTime(task.EndAt.String)
		if err != nil {
			return "", fmt.Errorf("invalid end_at [%d] '%s': %v", task.ID, task.EndAt.String, err)
		}
		endAt = endAt.Local()

		duration := minutesToText(endAt.Sub(startAt).Minutes())

		if isSameDate(startAt, endAt) {
			return fmt.Sprintf("%s - %s = %s", startAt.Format(dateTimeLayout), endAt.Format(timeLayout), duration), nil
		} else {
			return fmt.Sprintf("%s - %s = %s", startAt.Format(dateTimeLayout), endAt.Format(dateTimeLayout), duration), nil
		}
	} else {
		return fmt.Sprintf("%s - pending", startAt.Format(dateTimeLayout)), nil
	}
}

func isSameDate(t1, t2 time.Time) bool {
	if t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day() {
		return true
	}
	return false
}

func minutesToText(spanF float64) string {
	span := int(spanF)
	sign := ""
	if span < 0 {
		sign = "-"
		span = -span
	}

	hours := span / 60
	minutes := span % 60
	return fmt.Sprintf("%s%d:%02d", sign, hours, minutes)
}
