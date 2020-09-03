package payjputijl

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/payjp/payjp-go/v1"
)

var pay *payjp.Service

func init() {
	payjpSkTest := os.Getenv("PAYJP_SK_TEST")
	if payjpSkTest == "" {
		log.Println("failed to read env 'PAYJP_SK_TEST'")
	}
	pay = payjp.New(payjpSkTest, nil)
}

// Charge - 支払いをしてcharge_idを返す
func Charge(cardToken string, price, lessonID int64) (chargeID string, err error) {
	// 支払い
	charge, err := pay.Charge.Create(int(price), payjp.Charge{
		// 現在はjpyのみサポート
		Currency:    "jpy",
		CardToken:   cardToken,
		Capture:     true,
		Description: "charge for lesson",
		Metadata: map[string]string{
			"lesson_id": strconv.Itoa(int(lessonID)),
		},
	})
	log.Printf("payjp charge response: %+v\n", charge)
	if err != nil {
		log.Println("failed to charge, err:", err)
		return "", err
	}
	if charge.FailureMessage != "" {
		err := fmt.Errorf("payjp error: FailureCode: %v, FailureMessage: %v", charge.FailureCode, charge.FailureMessage)
		return chargeID, err
	}
	chargeID = charge.ID
	return chargeID, nil
}
