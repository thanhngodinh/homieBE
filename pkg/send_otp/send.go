package send_otp

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/sfreiberg/gotwilio"
)

const (
	accountSid   = "AC7729892a1c8f8af9cc45b5a7cbc9af9b"
	authToken    = "ee47313e4e1c562eb1554404c2c73fa7"
	twilioNumber = "+14849914441"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendOTP(phoneNumber, otp string) error {
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	message := "Your Homie OTP is: " + otp

	_, exception, err := twilio.SendSMS(twilioNumber, phoneNumber, message, "", "")
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	if exception != nil {
		return fmt.Errorf(exception.Message)
	}

	return nil
}
