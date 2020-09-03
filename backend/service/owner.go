package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-2020-b/dbutil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/repository"
)

type Owner struct {
	db *sqlx.DB
}

func NewOwner(db *sqlx.DB) *Owner {
	return &Owner{db}
}

func (o *Owner) Update(owner *model.Owner, bankAccount *model.BankAccount) (int64, error) {
	if err := dbutil.TXHandler(o.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateOwner(tx, owner)
		if err != nil {
			return err
		}
		_, err = repository.UpdateBankAccount(tx, bankAccount)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return 0, errors.Wrap(err, "failed owner and banck_account update transaction")
	}
	return owner.ID, nil
}

func (o *Owner) Create(owner *model.Owner, bankAccount *model.BankAccount) (int64, error) {
	var createdOwnerId int64
	if err := dbutil.TXHandler(o.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateOwner(tx, owner)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdOwnerId = id
		bankAccount.OwnerID = id
		_, err = repository.CreateBankAccount(tx, bankAccount)
		if err != nil {
			return err
		}
		createdOwnerId = id
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed to insert owner and bank account in the transaction")
	}
	return createdOwnerId, nil
}

func (o *Owner) GetByFirebaseUID(firebaseUID string) (*model.Owner, error) {
	var owner *model.Owner
	if err := dbutil.TXHandler(o.db, func(tx *sqlx.Tx) error {
		var err error
		owner, err = repository.GetOwnerByFirebaseID(tx, firebaseUID)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed to get owner by firebase_uid")
	}
	return owner, nil
}
