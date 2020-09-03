package zoom

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/voyagegroup/treasure-2020-b/model"
)

type ZoomAuthClient struct {
	ClientID            string
	ClientSecret        string
	ClientBasicToken    string
	BackendRedirectURI  string
	FrontendRedirectURI string
}

func NewZoomAuthClient(clientID string, clientSecret string, backendRedirectURI string, frontendRedirectURI string) *ZoomAuthClient {
	clientBasicToken := generateBasicToken(clientID, clientSecret)
	return &ZoomAuthClient{
		ClientID:            clientID,
		ClientSecret:        clientSecret,
		ClientBasicToken:    clientBasicToken,
		BackendRedirectURI:  backendRedirectURI,
		FrontendRedirectURI: frontendRedirectURI,
	}
}

func generateBasicToken(clientID string, clientSecret string) string {
	credential := []byte(clientID + ":" + clientSecret)
	token := base64.StdEncoding.EncodeToString(credential)

	return token
}

// https://marketplace.zoom.us/docs/guides/auth/oauth#step-2-request-access-token
func (z *ZoomAuthClient) OAuthReqAccessToken(redirectURI string, authorizationCode string, ownerID int64) (*model.ZoomToken, error) {
	log.Println("OAuthReqAccessToken:")

	client := &http.Client{}
	data := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {authorizationCode},
		"redirect_uri": {redirectURI},
	}
	log.Println("Request body", data)

	req, _ := http.NewRequest(
		"POST",
		"https://zoom.us/oauth/token",
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Authorization", "Basic "+z.ClientBasicToken)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var f interface{}
	json.Unmarshal(body, &f)
	respBody := f.(map[string]interface{})

	if resp.StatusCode != http.StatusOK {
		log.Println(respBody)
		errorReason := respBody["reason"].(string)
		return nil, errors.New(errorReason)
	}

	accessToken := respBody["access_token"].(string)
	refreshToken := respBody["refresh_token"].(string)
	zoomToken := &model.ZoomToken{
		AcccessToken: accessToken,
		RefreshToken: refreshToken,
		OwnerID:      ownerID,
	}
	log.Println("Zoom token", zoomToken)

	return zoomToken, nil
}

// https://marketplace.zoom.us/docs/guides/auth/oauth#refreshing
func (z *ZoomAuthClient) OAuthRefreshToken(zoomToken model.ZoomToken) (*model.ZoomToken, error) {
	log.Println("OAuthRefreshToken:")

	client := &http.Client{}
	data := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {zoomToken.RefreshToken},
	}
	log.Println("Request body", data)

	req, _ := http.NewRequest(
		"POST",
		"https://zoom.us/oauth/token",
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Authorization", "Basic "+z.ClientBasicToken)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to request zoom access token")
	}

	var f interface{}
	json.Unmarshal(body, &f)
	respBody := f.(map[string]interface{})

	newAccessToken := respBody["access_token"].(string)
	newRefreshToken := respBody["refresh_token"].(string)
	newZoomToken := &model.ZoomToken{
		OwnerID:      zoomToken.OwnerID,
		AcccessToken: newAccessToken,
		RefreshToken: newRefreshToken,
	}
	log.Println("New Zoom token", newZoomToken)

	return newZoomToken, nil
}
