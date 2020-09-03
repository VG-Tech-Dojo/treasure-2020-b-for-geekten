package model

import "time"

type Owner struct {
	ID              int64      `db:"id" json:"id"`
	FirstName       string     `db:"first_name" json:"first_name"`
	LastName        string     `db:"last_name" json:"last_name"`
	PostalNumber    string     `db:"postal_number" json:"postal_number"`
	Prefecture      string     `db:"prefecture" json:"prefecture"`
	City            string     `db:"city" json:"city"`
	Address         string     `db:"address" json:"address"`
	AddressOptional string     `db:"address_optional" json:"address_optional"`
	PhoneNumber     string     `db:"phone_number" json:"phone_number"`
	Email           string     `db:"email" json:"email"`
	FirebaseUID     string     `db:"firebase_uid" json:"firebase_uid"`
	CreatedAt       *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at"`
}
