package zoom

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/voyagegroup/treasure-2020-b/model"
)

func GenerateRandomPassword(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@-_*"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func GetZoomUserID(accessToken string) (string, error) {
	log.Println("GetZoomUserID:")

	client := &http.Client{}
	req, _ := http.NewRequest(
		"GET",
		"https://api.zoom.us/v2/users/me",
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(body)
		return "", errors.New("Failed to get zoom user info")
	}

	var f interface{}
	json.Unmarshal(body, &f)
	respBody := f.(map[string]interface{})

	zoomUserID := respBody["id"].(string)

	log.Println("Zoom user id", zoomUserID)

	return zoomUserID, nil
}

func (z *ZoomAuthClient) CreateMeeting(zoomID string, accessToken string, lesson *model.Lesson) (string, error) {
	log.Println("CreateMeeting:")

	durationMinutes := lesson.EndTime.Sub(*lesson.StartTime).Minutes()
	duration := strconv.Itoa(int(durationMinutes))

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	startTimeJST := lesson.StartTime.In(jst).Format(time.RFC3339)[:19]
	log.Println(startTimeJST)

	password, _ := GenerateRandomPassword(10)

	jsonStr := `{
  "topic": "パーソナルトレーニング",
  "type": 2,
  "start_time": "` + startTimeJST + `",
  "duration": ` + duration + `,
  "timezone": "Asia/Tokyo",
  "password": "` + password + `",
  "agenda": "パーソナルトレーニングの説明",
  "settings": {
    "join_before_host": true,
    "waiting_room": false
  }
}`
	log.Println("Request body", jsonStr)

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://api.zoom.us/v2/users/"+zoomID+"/meetings",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", errors.New("Failed to create zoom meeting")
	}

	var f interface{}
	json.Unmarshal(body, &f)
	respBody := f.(map[string]interface{})

	for k, v := range respBody {
		log.Printf("key=%s, type=%T, value=%v\n", k, v, v)
	}

	meetingURL := respBody["join_url"].(string)

	log.Println("Created meeting id", meetingURL)

	return meetingURL, nil
}
