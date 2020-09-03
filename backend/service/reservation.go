package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-2020-b/dbutil"
	"github.com/voyagegroup/treasure-2020-b/model"
	payjputijl "github.com/voyagegroup/treasure-2020-b/payjputil"
	"github.com/voyagegroup/treasure-2020-b/repository"
	"github.com/voyagegroup/treasure-2020-b/sendgridutil"
)

type Reservation struct {
	db *sqlx.DB
}

func NewReservation(db *sqlx.DB) *Reservation {
	return &Reservation{db}
}

func (l *Reservation) Create(cardToken string, reservation *model.Reservation) error {
	if err := dbutil.TXHandler(l.db, func(tx *sqlx.Tx) error {

		chargeID, err := payjputijl.Charge(cardToken, reservation.PaidPrice, reservation.LessonID)
		if err != nil {
			return err
		}
		lesson, err := repository.GetLessonByID(tx, reservation.LessonID)
		if err != nil {
			return err
		}
		owner, err := repository.GetOwnerByID(tx, lesson.OwnerID)
		if err != nil {
			return err
		}

		meetingURL := lesson.MeetingID

		if err = sendgridutil.SendMailClient(reservation, lesson, owner, meetingURL); err != nil {
			return err
		}
		if err = sendgridutil.SendMailOwner(reservation, lesson, owner, meetingURL); err != nil {
			return err
		}

		reservation.ChargeID = chargeID
		_, err = repository.CreateReservation(tx, reservation)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed to insert reservation in the transaction")
	}
	return nil
}
