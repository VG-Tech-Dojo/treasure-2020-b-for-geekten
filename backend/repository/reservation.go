package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/model"
)

func CreateReservation(db *sqlx.Tx, r *model.Reservation) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
INSERT INTO reservation (lesson_id, charge_id, paid_price, first_name, last_name, email) 
VALUES (?, ?, ?, ?, ?, ?)
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
		r.LessonID,
		r.ChargeID,
		r.PaidPrice,
		r.FirstName,
		r.LastName,
		r.Email,
	)
}
