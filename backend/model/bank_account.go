package model

type BankAccount struct {
	ID                    int64  `db:"id" json:"id"`
	OwnerID               int64  `db:"owner_id" json:"owner_id"`
	BankAccountNumber     string `db:"bank_account_number" json:"bank_account_number"`
	BankBranchCode        string `db:"bank_branch_code" json:"bank_branch_code"`
	BankCode              string `db:"bank_code" json:"bank_code"`
	BankAccountHolderName string `db:"bank_account_holder_name" json:"bank_account_holder_name"`
	BankAccountType       string `db:"bank_account_type" json:"bank_account_type"`
}
