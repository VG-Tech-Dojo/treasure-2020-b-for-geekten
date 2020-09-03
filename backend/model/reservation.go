package model

import "time"

type Reservation struct {
	ID        int64  `db:"id" json:"id"`
	LessonID  int64  `db:"lesson_id" json:"lesson_id"`
	ChargeID  string `db:"charge_id" json:"charge_id"`
	PaidPrice int64  `db:"paid_price" json:"paid_price"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	CreatedAt *time.Time
}
