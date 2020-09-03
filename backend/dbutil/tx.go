package dbutil

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
)

// TXHandler is handler for working with transaction.
// This is wrapper function for commit and rollback.
func TXHandler(db *sqlx.DB, f func(*sqlx.Tx) error) error {
	var err error
	tx, err := db.Beginx()
	if err != nil {
		return errors.Wrap(err, "start transaction failed")
	}
	defer func() {
		if err := recover(); err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				log.Fatalf("rollback failed: %v", rollBackErr)
			}
			log.Print("Rollback operation")
			return
		}
	}()
	if err := f(tx); err != nil {
		return errors.Wrap(err, "transaction: operation failed")
	}
	return nil
}
