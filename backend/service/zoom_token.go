package service

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-2020-b/dbutil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/repository"
)

type ZoomToken struct {
	db *sqlx.DB
}

func NewZoomToken(db *sqlx.DB) *ZoomToken {
	return &ZoomToken{db}
}

func (z *ZoomToken) Create(zoomToken *model.ZoomToken) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(z.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateZoomToken(tx, zoomToken)
		if err != nil {
			log.Println("ERRRRRRRRRRRRR 1")
			return err
		}
		if err := tx.Commit(); err != nil {
			log.Println("ERRRRRRRRRRRRR 2")
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			log.Println("ERRRRRRRRRRRRR 3")
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed zoom_token insert transaction")
	}
	return createdId, nil
}
