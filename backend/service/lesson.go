package service

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/voyagegroup/treasure-2020-b/dbutil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/repository"
	"github.com/voyagegroup/treasure-2020-b/zoom"
)

type Lesson struct {
	db         *sqlx.DB
	authClient *zoom.ZoomAuthClient
}

func NewLesson(db *sqlx.DB, authClient *zoom.ZoomAuthClient) *Lesson {
	return &Lesson{
		db:         db,
		authClient: authClient,
	}
}

func (l *Lesson) Create(firebaseUID string, lesson *model.Lesson) (int64, error) {
	var createdLessonId int64
	if err := dbutil.TXHandler(l.db, func(tx *sqlx.Tx) error {

		//TODO zoomのmeetingを作成してlesson.MeetingIDに代入する

		owner, err := repository.GetOwnerByFirebaseID(tx, firebaseUID)
		if err != nil {
			log.Println(err)
			return err
		}
		lesson.OwnerID = owner.ID

		zoomToken, err := repository.GetZoomTokenByOwnerID(l.db, owner.ID)
		if err != nil {
			log.Println(err)
			return err
		}

		zoomToken, err = l.authClient.OAuthRefreshToken(*zoomToken)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(zoomToken)

		_, err = repository.CreateZoomToken(tx, zoomToken)
		if err != nil {
			log.Println(err)
			return err
		}

		zoomUserID, err := zoom.GetZoomUserID(zoomToken.AcccessToken)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(zoomUserID)

		meetingID, err := l.authClient.CreateMeeting(zoomUserID, zoomToken.AcccessToken, lesson)
		if err != nil {
			log.Println(err)
			return err
		}

		lesson.MeetingID = meetingID

		result, err := repository.CreateLesson(tx, lesson)
		log.Println("insert lesson result:", result)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		lesson.ID = id
		createdLessonId = id

		return err
	}); err != nil {
		log.Println(err)
		return 0, errors.Wrap(err, "failed to insert lesson in the transaction")
	}
	return createdLessonId, nil
}

func (l *Lesson) GetByID(lessonID int64) (*model.Lesson, error) {
	var lesson *model.Lesson
	if err := dbutil.TXHandler(l.db, func(tx *sqlx.Tx) error {
		var err error
		lesson, err = repository.GetLessonByID(tx, lessonID)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed to select lesson in the transaction")
	}
	return lesson, nil
}

func (l *Lesson) GetAll(firebaseUID string) ([]model.Lesson, error) {
	var lessons []model.Lesson
	if err := dbutil.TXHandler(l.db, func(tx *sqlx.Tx) error {
		owner, err := repository.GetOwnerByFirebaseID(tx, firebaseUID)
		if err != nil {
			return err
		}
		lessons, err = repository.AllLesson(l.db, owner.ID)
		if err == sql.ErrNoRows {
			return nil
		} else if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed to select lessons in the transaction")
	}
	return lessons, nil
}

func (l *Lesson) GetAllByOwnerID(ownerID int64) ([]model.Lesson, error) {
	var lessons []model.Lesson
	if err := dbutil.TXHandler(l.db, func(tx *sqlx.Tx) error {
		var err error
		lessons, err = repository.AllLesson(l.db, ownerID)
		if err == sql.ErrNoRows {
			return nil
		} else if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, errors.Wrap(err, "failed to select lessons in the transaction")
	}
	return lessons, nil
}
