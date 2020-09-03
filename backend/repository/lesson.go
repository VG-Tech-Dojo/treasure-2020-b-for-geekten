package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-2020-b/model"
)

func CreateLesson(db *sqlx.Tx, l *model.Lesson) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
INSERT INTO lesson (owner_id, start_time, end_time, meeting_id, price) 
VALUES (?, ?, ?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return stmt.Exec(
		l.OwnerID,
		l.StartTime,
		l.EndTime,
		l.MeetingID,
		l.Price,
	)
}

func AllLesson(db *sqlx.DB, ownerID int64) ([]model.Lesson, error) {
	lessons := make([]model.Lesson, 0)
	if err := db.Select(&lessons, `SELECT id, owner_id, start_time, end_time, meeting_id, price 
FROM lesson
WHERE owner_id = ?
`, ownerID); err != nil {
		return nil, err
	}
	return lessons, nil
}

func GetLessonByID(db *sqlx.Tx, id int64) (*model.Lesson, error) {
	l := model.Lesson{}
	if err := db.Get(&l, `
SELECT id, owner_id, start_time, end_time, meeting_id, price 
FROM lesson WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &l, nil
}
