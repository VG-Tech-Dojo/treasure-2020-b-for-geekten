package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/model"
)

func CreateOwner(db *sqlx.Tx, o *model.Owner) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
INSERT INTO owner (first_name, last_name, postal_number, prefecture, city, address, address_optional, phone_number, email, firebase_uid) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
		o.FirstName,
		o.LastName,
		o.PostalNumber,
		o.Prefecture,
		o.City,
		o.Address,
		o.AddressOptional,
		o.PhoneNumber,
		o.Email,
		o.FirebaseUID,
	)
}

func UpdateOwner(db *sqlx.Tx, o *model.Owner) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
UPDATE owner
SET first_name = ?, last_name = ?, postal_number = ?, prefecture = ?, city = ?, address = ?, address_optional = ?, phone_number = ?, email = ?, firebase_uid = ?
WHERE id = ?
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
		o.FirstName,
		o.LastName,
		o.PostalNumber,
		o.Prefecture,
		o.City,
		o.Address,
		o.AddressOptional,
		o.PhoneNumber,
		o.Email,
		o.FirebaseUID,
		o.ID,
	)
}

func GetOwnerByID(db *sqlx.Tx, id int64) (*model.Owner, error) {
	o := model.Owner{}
	if err := db.Get(&o, `
SELECT first_name, last_name, postal_number, prefecture, city, address, address_optional, phone_number, email, firebase_uid 
FROM owner WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &o, nil
}

func GetOwnerByFirebaseID(db *sqlx.Tx, firebaseUID string) (*model.Owner, error) {
	o := model.Owner{}
	if err := db.Get(&o, `
SELECT id, first_name, last_name, postal_number, prefecture, city, address, address_optional, phone_number, email 
FROM owner WHERE firebase_uid = ?
`, firebaseUID); err != nil {
		return nil, err
	}
	return &o, nil
}
