package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/voyagegroup/treasure-2020-b/httputil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/service"
)

type Owner struct {
	db *sqlx.DB
}

func NewOwner(db *sqlx.DB) *Owner {
	return &Owner{db: db}
}

func (o *Owner) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ownerRequestReceiver := &OwnerRequestReceiver{}
	if err := json.NewDecoder(r.Body).Decode(&ownerRequestReceiver); err != nil {
		return http.StatusBadRequest, nil, err
	}
	newOwner, newBankAccount, err := ownerRequestReceiver.Divide()
	if err != nil {
		return http.StatusUnprocessableEntity, nil, err
	}

	ownerService := service.NewOwner(o.db)
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	newOwner.FirebaseUID = user.FirebaseUID
	ownerID, err := ownerService.Create(newOwner, newBankAccount)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newOwner.ID = ownerID
	return http.StatusCreated, newOwner, nil
}

func (o *Owner) Update(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)

	// GET parameter /owner/:id
	ownerID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("owner id is invalid")
	}

	ownerRequestReceiver := &OwnerRequestReceiver{}
	if err := json.NewDecoder(r.Body).Decode(&ownerRequestReceiver); err != nil {
		return http.StatusBadRequest, nil, err
	}

	owner, bankAccount, err := ownerRequestReceiver.Divide()
	if err != nil {
		return http.StatusUnprocessableEntity, nil, err
	}

	owner.ID = ownerID

	ownerService := service.NewOwner(o.db)
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	owner.FirebaseUID = user.FirebaseUID
	_, err = ownerService.Update(owner, bankAccount)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, owner, nil
}

func (o *Owner) GetMe(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ownerService := service.NewOwner(o.db)
	owner, err := ownerService.GetByFirebaseUID(user.FirebaseUID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, owner, nil
}

type OwnerRequestReceiver struct {
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	PostalNumber          string `json:"postal_number"`
	Prefecture            string `json:"prefecture"`
	City                  string `json:"city"`
	Address               string `json:"address"`
	AddressOptional       string `json:"address_optional"`
	PhoneNumber           string `json:"phone_number"`
	Email                 string `json:"email"`
	BankAccountNumber     string `json:"bank_account_number"`
	BankBranchCode        string `json:"bank_branch_code"`
	BankCode              string `json:"bank_code"`
	BankAccountHolderName string `json:"bank_account_holder_name"`
	BankAccountType       string `json:"bank_account_type"`
}

func (r *OwnerRequestReceiver) Divide() (*model.Owner, *model.BankAccount, error) {
	newOwner := &model.Owner{
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		PostalNumber:    r.PostalNumber,
		Prefecture:      r.Prefecture,
		City:            r.City,
		Address:         r.Address,
		AddressOptional: r.AddressOptional,
		PhoneNumber:     r.PhoneNumber,
		Email:           r.Email,
	}
	newBankAccount := &model.BankAccount{
		BankAccountNumber:     r.BankAccountNumber,
		BankBranchCode:        r.BankBranchCode,
		BankCode:              r.BankCode,
		BankAccountHolderName: r.BankAccountHolderName,
		BankAccountType:       r.BankAccountType,
	}
	err := r.Validate()
	return newOwner, newBankAccount, err
}

func (r OwnerRequestReceiver) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.FirstName, validation.Required),
		validation.Field(&r.LastName, validation.Required),
		validation.Field(&r.PostalNumber, validation.Required),
		validation.Field(&r.Prefecture, validation.Required),
		validation.Field(&r.City, validation.Required),
		validation.Field(&r.Address, validation.Required),
		validation.Field(&r.PhoneNumber, validation.Required),
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.BankAccountNumber, validation.Required),
		validation.Field(&r.BankBranchCode, validation.Required),
		validation.Field(&r.BankCode, validation.Required),
		validation.Field(&r.BankAccountHolderName, validation.Required),
		validation.Field(&r.BankAccountType, validation.Required),
	)
}
