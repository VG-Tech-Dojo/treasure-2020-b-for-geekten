package sendgridutil

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/voyagegroup/treasure-2020-b/model"
	"log"
	"os"
	"strings"
	"time"
)

var serviceEmail string
var sendgridAPIKey string
var loc *time.Location

const (
	mailFrom        = "SlimLine運営"
	mailTitleClient = "【SlimLine】ご予約ありがとうございます"
	mailTitleOwner  = "【SlimLine】レッスンの予約が成立しました"

	dateFormat       = "1/2"
	hourMinuteFormat = "15:04"
	location         = "Asia/Tokyo"
)

func init() {
	serviceEmail = os.Getenv("SENDGRID_SERVICE_EMAIL")
	if serviceEmail == "" {
		log.Println("failed to read env 'SENDGRID_SERVICE_EMAIL'")
	}
	sendgridAPIKey = os.Getenv("SENDGRID_APIKEY")
	if sendgridAPIKey == "" {
		log.Println("failed to read env 'SENDGRID_APIKEY'")
	}
	var err error
	loc, err = time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
}

func SendMailClient(reservation *model.Reservation, lesson *model.Lesson, owner *model.Owner, meetingURL string) error {
	from := mail.NewEmail(mailFrom, serviceEmail)
	subject := mailTitleClient
	to := mail.NewEmail(reservation.LastName+reservation.FirstName+"様", reservation.Email)
	htmlContent := fmt.Sprintf(`%s 様<br>
<br>
こんにちは！<br>
パーソナルトレーナーの%sです。<br>
この度はパーソナルトレーニングのご予約、ありがとうございます。<br>
<br>
--------- ご予約内容 ---------<br>
講師：%s<br>
受講日時：%s %s ~ %s<br>
料金：￥%d<br>
Zoom URL：<a href="%s">%s</a><br>
-----------------------------<br>
<br>
お時間になりましたら上記のZoom URLにアクセスして受講を開始してください。<br>
<br>
よろしくお願いいたします。<br>
<br>
オンラインレッスン支援サービス SlimLine <br>
あなたのレッスンを、気軽にオンラインで。<br>
<a href="https://group-b.treasure2020.dojo-voyage.net">https://group-b.treasure2020.dojo-voyage.net</a><br>
`,
		reservation.LastName+reservation.FirstName,
		owner.LastName+owner.FirstName,
		owner.LastName+owner.FirstName,
		lesson.StartTime.In(loc).Format(dateFormat),
		lesson.StartTime.In(loc).Format(hourMinuteFormat),
		lesson.EndTime.In(loc).Format(hourMinuteFormat),
		reservation.PaidPrice,
		meetingURL,
		meetingURL,
	)
	plainTextContent := strings.ReplaceAll(htmlContent, "<br>", "\n")
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendgridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	log.Printf("sendgrid client response.StatusCode: %v", response.StatusCode)
	log.Printf("sendgrid client response.Headers: %v", response.Headers)
	log.Printf("sendgrid client response.Body: %v", response.Body)
	return nil
}

func SendMailOwner(reservation *model.Reservation, lesson *model.Lesson, owner *model.Owner, meetingURL string) error {
	from := mail.NewEmail(mailFrom, serviceEmail)
	subject := mailTitleOwner
	to := mail.NewEmail(owner.LastName+owner.FirstName+"様", owner.Email)
	htmlContent := fmt.Sprintf(`%s 様<br>
<br>
SlimLine運営です。<br>
レッスンが新たに予約されたことをお知らせいたします。<br>
<br>
--------- 予約内容 ---------<br>
名前：%s<br>
受講日時：%s %s ~ %s<br>
料金：￥%d<br>
Zoom URL：<a href="%s">%s</a><br>
----------------------------<br>
<br>
お時間になりましたら上記のZoom URLにアクセスして<br>
パーソナルトレーニングを開始してください。<br>
<br>
よろしくお願いいたします。<br>
<br>
オンラインレッスン支援サービス SlimLine <br>
あなたのレッスンを、気軽にオンラインで。<br>
<a href="https://group-b.treasure2020.dojo-voyage.net">https://group-b.treasure2020.dojo-voyage.net</a><br>
`,
		owner.LastName+owner.FirstName,
		reservation.LastName+reservation.FirstName,
		lesson.StartTime.In(loc).Format(dateFormat),
		lesson.StartTime.In(loc).Format(hourMinuteFormat),
		lesson.EndTime.In(loc).Format(hourMinuteFormat),
		reservation.PaidPrice,
		meetingURL,
		meetingURL,
	)
	plainTextContent := strings.ReplaceAll(htmlContent, "<br>", "\n")
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendgridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	log.Printf("sendgrid client response.StatusCode: %v", response.StatusCode)
	log.Printf("sendgrid client response.Headers: %v", response.Headers)
	log.Printf("sendgrid client response.Body: %v", response.Body)
	return nil
}
