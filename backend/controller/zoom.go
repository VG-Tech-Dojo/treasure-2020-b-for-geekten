package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/dbutil"
	"github.com/voyagegroup/treasure-2020-b/model"
	"github.com/voyagegroup/treasure-2020-b/repository"
	"github.com/voyagegroup/treasure-2020-b/service"
	"github.com/voyagegroup/treasure-2020-b/zoom"
)

type Zoom struct {
	db         *sqlx.DB
	authClient *zoom.ZoomAuthClient
}

func NewZoom(db *sqlx.DB, authClient *zoom.ZoomAuthClient) *Zoom {
	return &Zoom{
		db:         db,
		authClient: authClient,
	}
}

func (z *Zoom) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	log.Println(vars)

	firebaseUID := vars["firebase_uid"]

	var err error
	var owner *model.Owner
	if err := dbutil.TXHandler(z.db, func(tx *sqlx.Tx) error {
		owner, err = repository.GetOwnerByFirebaseID(tx, firebaseUID)
		if err != nil {
			return err
		}
		return err
	}); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	log.Println("Owner", owner)

	authorizationCode := r.FormValue("code")
	if authorizationCode == "" {
		return http.StatusBadRequest, nil, errors.New("Authorization code is missing")
	}

	log.Println("ZoomOAuthGETHandler:")
	log.Println("Authorization code", authorizationCode)

	reqAccessTokenRedirectURI := z.authClient.BackendRedirectURI + r.URL.Path
	log.Println("Redirect URL", reqAccessTokenRedirectURI)

	zoomToken, err := z.authClient.OAuthReqAccessToken(reqAccessTokenRedirectURI, authorizationCode, owner.ID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	zoomToken.OwnerID = owner.ID
	log.Println(zoomToken)

	zoomTokenService := service.NewZoomToken(z.db)
	id, err := zoomTokenService.Create(zoomToken)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, nil, err
	}
	zoomToken.ID = id

	http.Redirect(w, r, z.authClient.FrontendRedirectURI, 301)

	return http.StatusCreated, nil, nil
}
