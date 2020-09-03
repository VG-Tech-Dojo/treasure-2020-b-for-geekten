package controller

import (
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/repository"
	"github.com/voyagegroup/treasure-2020-b/service"
	"net/http"
)

type Room struct {
	db *sqlx.DB
}

func NewRoom(db *sqlx.DB) *Room {
	return &Room{db: db}
}

func (a *Room) Index(_ http.ResponseWriter, _ *http.Request) (int, interface{}, error) {
	rooms, err := repository.AllRoom(a.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, rooms, nil
}

func (a *Room) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newRoom := &model.Room{}
	if err := json.NewDecoder(r.Body).Decode(&newRoom); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if newRoom.Title == "" {
		return http.StatusUnprocessableEntity, nil, errors.New("required parameter is missing")
	}

	roomService := service.NewRoom(a.db)
	id, err := roomService.Create(newRoom)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newRoom.ID = id
	return http.StatusCreated, newRoom, nil
}

func (a *Room) Update(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusNotFound, nil, nil
}

func (a *Room) Destroy(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusNotFound, nil, nil
}
