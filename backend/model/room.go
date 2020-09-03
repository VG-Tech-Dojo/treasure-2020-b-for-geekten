package model

type Room struct {
	ID    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}
