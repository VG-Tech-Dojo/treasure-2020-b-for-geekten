package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/service"
)

type ReservationJSONReceiver struct {
	PaidPrice int64  `db:"paid_price" json:"paid_price"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	CardToken string `db:"card_token" json:"card_token"`
}

func ReservationJSONToModel(r *ReservationJSONReceiver) *model.Reservation {
	return &model.Reservation{
		PaidPrice: r.PaidPrice,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
	}
}

func ReservationModelToJSON(r *model.Reservation) *ReservationJSONReceiver {
	return &ReservationJSONReceiver{
		PaidPrice: r.PaidPrice,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
	}
}

type Reservation struct {
	db *sqlx.DB
}

func NewReservation(db *sqlx.DB) *Reservation {
	return &Reservation{db: db}
}

func (l *Reservation) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	lessonID, err := strconv.ParseInt(vars["lesson_id"], 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("lesson id is invalid")
	}

	reservationJSONReceiver := &ReservationJSONReceiver{}
	if err := json.NewDecoder(r.Body).Decode(&reservationJSONReceiver); err != nil {
		return http.StatusBadRequest, nil, err
	}
	reservation := ReservationJSONToModel(reservationJSONReceiver)
	reservation.LessonID = lessonID

	reservationService := service.NewReservation(l.db)
	err = reservationService.Create(reservationJSONReceiver.CardToken, reservation)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	reservationJSONReceiver = ReservationModelToJSON(reservation)

	return http.StatusCreated, reservationJSONReceiver, nil
}
