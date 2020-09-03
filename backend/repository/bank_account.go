package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/model"
)

func CreateBankAccount(db *sqlx.Tx, b *model.BankAccount) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
INSERT INTO bank_account (owner_id, bank_account_number, bank_branch_code, bank_code, bank_account_holder_name, bank_account_type) 
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
		b.OwnerID,
		b.BankAccountNumber,
		b.BankBranchCode,
		b.BankCode,
		b.BankAccountHolderName,
		b.BankAccountType,
	)
}

func UpdateBankAccount(db *sqlx.Tx, b *model.BankAccount) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
UPDATE bank_account
SET owner_id = ?, bank_account_number = ?, bank_branch_code = ?, bank_code = ?, bank_account_holder_name = ?, bank_account_type = ?
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
		b.OwnerID,
		b.BankAccountNumber,
		b.BankBranchCode,
		b.BankCode,
		b.BankAccountHolderName,
		b.BankAccountType,
	)
}
