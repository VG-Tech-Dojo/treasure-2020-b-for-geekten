package model

import (
	"time"
)

type Lesson struct {
	ID        int64      `db:"id" json:"id"`
	OwnerID   int64      `db:"owner_id" json:"owner_id"`
	StartTime *time.Time `db:"start_time" json:"start_time"`
	EndTime   *time.Time `db:"end_time" json:"end_time"`
	MeetingID string     `db:"meeting_id" json:"meeting_id"`
	Price     int64      `db:"price" json:"price"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
