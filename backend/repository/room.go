package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/model"
)

func AllRoom(db *sqlx.DB) ([]model.Room, error) {
	a := make([]model.Room, 0)
	if err := db.Select(&a, `SELECT id, title FROM room`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindRoom(db *sqlx.DB, id int64) (*model.Room, error) {
	a := model.Room{}
	if err := db.Get(&a, `
SELECT id, title FROM room WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateRoom(db *sqlx.Tx, a *model.Room) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
INSERT INTO room (title) VALUES (?)
`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return stmt.Exec(a.Title)
}
