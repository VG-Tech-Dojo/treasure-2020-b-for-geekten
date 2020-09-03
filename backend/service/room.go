package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/voyagegroup/treasure-2020-b/dbutil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/repository"
)

type Room struct {
	db *sqlx.DB
}

func NewRoom(db *sqlx.DB) *Room {
	return &Room{db}
}

func (a *Room) Create(room *model.Room) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateRoom(tx, room)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed room insert transaction")
	}
	return createdId, nil
}
