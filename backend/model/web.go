package model

import "time"

type Web struct {
	ID        int64      `db:"id" json:"id"`
	OwnerID   int64      `db:"owner_id" json:"owner_id"`
	Title     string     `db:"title" json:"title"`
	Profile   string     `db:"profile" json:"profile"`
	Theme     string     `db:"theme" json:"theme"`
	Content   string     `db:"content" json:"content"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
