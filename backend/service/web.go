package service

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-2020-b/dbutil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/repository"
)

type Web struct {
	db *sqlx.DB
}

func NewWeb(db *sqlx.DB) *Web {
	return &Web{db}
}

func (w *Web) Create(firebaseUID string, web *model.Web) (*model.Web, error) {
	if err := dbutil.TXHandler(w.db, func(tx *sqlx.Tx) error {
		owner, err := repository.GetOwnerByFirebaseID(tx, firebaseUID)
		if err != nil {
			return err
		}
		web.OwnerID = owner.ID
		result, err := repository.CreateWeb(tx, web)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		web.ID = id

		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed to insert web in the transaction")
	}
	return web, nil
}

func (w *Web) GetAll(firebaseUID string) ([]model.Web, error) {
	var webs []model.Web
	if err := dbutil.TXHandler(w.db, func(tx *sqlx.Tx) error {
		owner, err := repository.GetOwnerByFirebaseID(tx, firebaseUID)
		if err != nil {
			return err
		}
		webs, err = repository.AllWeb(w.db, owner.ID)
		if err != sql.ErrNoRows{
			return nil
		}else if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed to select webs in the transaction")
	}
	return webs, nil
}

func (w *Web) GetAllByOwnerID(ownerID int64) ([]model.Web, error) {
	var webs []model.Web
	if err := dbutil.TXHandler(w.db, func(tx *sqlx.Tx) error {
		var err error
		webs, err = repository.AllWeb(w.db, ownerID)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed to select webs in the transaction")
	}
	return webs, nil
}
