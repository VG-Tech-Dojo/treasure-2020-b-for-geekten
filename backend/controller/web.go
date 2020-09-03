package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/httputil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/repository"
	"github.com/voyagegroup/treasure-2020-b/service"
)

type Web struct {
	db *sqlx.DB
}

func NewWeb(db *sqlx.DB) *Web {
	return &Web{db: db}
}

func (w *Web) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	web := &model.Web{}
	if err := json.NewDecoder(r.Body).Decode(web); err != nil {
		return http.StatusBadRequest, nil, err
	}

	webService := service.NewWeb(w.db)
	web, err = webService.Create(user.FirebaseUID, web)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, web, nil
}

func (w *Web) GetAll(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	webService := service.NewWeb(w.db)
	webs, err := webService.GetAll(user.FirebaseUID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, webs, nil
}

func (w *Web) GetAllByOwnerID(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	ownerID, err := strconv.ParseInt(vars["owner_id"], 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("owner id is invalid")
	}

	webService := service.NewWeb(w.db)
	webs, err := webService.GetAllByOwnerID(ownerID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if len(webs) == 0 {
		return http.StatusNotFound, nil, errors.New("Web site is not created")
	}

	return http.StatusOK, webs, nil
}

func (w *Web) Get(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	webID, err := strconv.ParseInt(vars["web_id"], 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("web id is invalid")
	}
	web, err := repository.GetWebByID(w.db, webID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, web, nil
}
