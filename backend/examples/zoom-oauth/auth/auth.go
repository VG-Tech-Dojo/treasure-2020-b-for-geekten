package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/tokoroten-lab/oauth-test/model"
)

func generateBasicToken(clientID string, clientSecret string) string {
	credential := []byte(clientID + ":" + clientSecret)
	token := base64.StdEncoding.EncodeToString(credential)

	return token
}

// https://marketplace.zoom.us/docs/guides/auth/oauth#step-2-request-access-token
func OAuthReqAccessToken(clientID string, clientSecret string, redirectURI string, authorizationCode string, ownerID int64) (*model.ZoomToken, error) {
	fmt.Println("OAuthReqAccessToken:")

	basicToken := generateBasicToken(clientID, clientSecret)
	fmt.Println("Basic token", basicToken)

	client := &http.Client{}
	data := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {authorizationCode},
		"redirect_uri": {redirectURI},
	}
	fmt.Println("Request body", data)

	req, _ := http.NewRequest(
		"POST",
		"https://zoom.us/oauth/token",
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Authorization", "Basic "+basicToken)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var f interface{}
	json.Unmarshal(body, &f)
	respBody := f.(map[string]interface{})

	if resp.StatusCode != http.StatusOK {
		fmt.Println(respBody)
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
	println("Zoom token", zoomToken)

	return zoomToken, nil
}

// https://marketplace.zoom.us/docs/guides/auth/oauth#refreshing
func OAuthRefreshToken(clientID string, clientSecret string, zoomToken model.ZoomToken) (*model.ZoomToken, error) {
	fmt.Println("OAuthRefreshToken:")

	basicToken := generateBasicToken(clientID, clientSecret)
	fmt.Println("Basic token", basicToken)

	client := &http.Client{}
	data := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {zoomToken.RefreshToken},
	}
	fmt.Println("Request body", data)

	req, _ := http.NewRequest(
		"POST",
		"https://zoom.us/oauth/token",
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Authorization", "Basic "+basicToken)

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
	println("New Zoom token", newZoomToken)

	return newZoomToken, nil
}
