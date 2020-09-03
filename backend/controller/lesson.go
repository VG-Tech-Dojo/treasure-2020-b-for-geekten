package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-2020-b/httputil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/service"
	"github.com/voyagegroup/treasure-2020-b/zoom"
)

type LessonJSONReceiver struct {
	ID        int64  `json:"id"`
	OwnerID   int64  `json:"owner_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	MeetingID string `json:"meeting_id"`
	Price     int64  `json:"price"`
}

func LessonJSONToModel(l *LessonJSONReceiver) *model.Lesson {
	startTime := time.Unix(l.StartTime, 0)
	endTime := time.Unix(l.EndTime, 0)
	return &model.Lesson{
		ID:        l.ID,
		OwnerID:   l.OwnerID,
		StartTime: &startTime,
		EndTime:   &endTime,
		MeetingID: l.MeetingID,
		Price:     l.Price,
	}
}

func LessonModelToJSON(l *model.Lesson) *LessonJSONReceiver {
	return &LessonJSONReceiver{
		ID:        l.ID,
		OwnerID:   l.OwnerID,
		StartTime: l.StartTime.Unix(),
		EndTime:   l.EndTime.Unix(),
		MeetingID: l.MeetingID,
		Price:     l.Price,
	}
}

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

func (l *Lesson) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	lessonJSONReceiver := &LessonJSONReceiver{}
	if err := json.NewDecoder(r.Body).Decode(&lessonJSONReceiver); err != nil {
		return http.StatusBadRequest, nil, err
	}
	lesson := LessonJSONToModel(lessonJSONReceiver)

	lessonService := service.NewLesson(l.db, l.authClient)
	lessonID, err := lessonService.Create(user.FirebaseUID, lesson)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	lesson.ID = lessonID

	lessonJSONReceiver = LessonModelToJSON(lesson)

	return http.StatusCreated, lessonJSONReceiver, nil
}

func (l *Lesson) GetByID(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	lessonID, err := strconv.ParseInt(vars["lesson_id"], 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("lesson id is invalid")
	}

	lessonService := service.NewLesson(l.db, l.authClient)
	lesson, err := lessonService.GetByID(lessonID)
	if err != nil {
		return http.StatusNotFound, nil, errors.New("lesson not found")
	}

	lessonJSONReceiver := *LessonModelToJSON(lesson)

	return http.StatusOK, lessonJSONReceiver, nil
}

func (l *Lesson) GetAll(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	lessonService := service.NewLesson(l.db, l.authClient)
	lessons, err := lessonService.GetAll(user.FirebaseUID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	lessonJSONReceivers := make([]LessonJSONReceiver, 0, len(lessons))
	for _, lesson := range lessons {
		lessonJSONReceivers = append(lessonJSONReceivers, *LessonModelToJSON(&lesson))
	}

	return http.StatusOK, lessonJSONReceivers, nil
}

func (l *Lesson) GetAllByOwnerID(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	ownerID, err := strconv.ParseInt(vars["owner_id"], 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("owner id is invalid")
	}
	lessonService := service.NewLesson(l.db, l.authClient)
	lessons, err := lessonService.GetAllByOwnerID(ownerID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	lessonJSONReceivers := make([]LessonJSONReceiver, 0, len(lessons))
	for _, lesson := range lessons {
		lessonJSONReceivers = append(lessonJSONReceivers, *LessonModelToJSON(&lesson))
	}

	return http.StatusOK, lessonJSONReceivers, nil
}
