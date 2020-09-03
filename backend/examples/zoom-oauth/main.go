package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tokoroten-lab/oauth-test/auth"
	"github.com/tokoroten-lab/oauth-test/model"
)

var zoomClientID string
var zoomClientSecret string
var zoomRedirectURI string

var zoomToken model.ZoomToken

func ZoomOAuthPOSTHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

	ownerID, err := strconv.ParseInt(vars["owner_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("owner_id", ownerID)

	location := zoomRedirectURI + fmt.Sprintf("/owner/%d/zoom_auth", ownerID)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", location)
	w.WriteHeader(http.StatusOK)
}

func ZoomOAuthGETHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

	ownerID, err := strconv.ParseInt(vars["owner_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authorizationCode := r.FormValue("code")
	if authorizationCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("ZoomOAuthGETHandler:")
	fmt.Println("Authorization code", authorizationCode)

	reqAccessTokenRedirectURI := zoomRedirectURI + r.URL.Path
	fmt.Println("Redirect URL", reqAccessTokenRedirectURI)

	zoomToken, err := auth.OAuthReqAccessToken(zoomClientID, zoomClientSecret, reqAccessTokenRedirectURI, authorizationCode, ownerID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(zoomToken)

	refreshToken, err := auth.OAuthRefreshToken(zoomClientID, zoomClientSecret, *zoomToken)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(refreshToken)

	zoomToken = refreshToken
}

func main() {
	flag.StringVar(&zoomClientID, "zoomClientID", "", "Client ID (Zoom app)")
	flag.StringVar(&zoomClientSecret, "zoomClientSecret", "", "Client Secret (Zoom app)")
	flag.StringVar(&zoomRedirectURI, "zoomRedirectURI", "", "Redirect URI (Zoom app)")
	flag.Parse()

	fmt.Println("Cient ID", zoomClientID)
	fmt.Println("Client Secret", zoomClientSecret)
	fmt.Println("Redirect URI", zoomRedirectURI)

	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path("/owner/{owner_id}/zoom_auth").HandlerFunc(ZoomOAuthPOSTHandler)
	r.Methods(http.MethodGet).Path("/owner/{owner_id}/zoom_auth").HandlerFunc(ZoomOAuthGETHandler)
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
