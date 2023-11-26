package send_email

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/sfreiberg/gotwilio"
)

const (
	accountSid   = "AC7729892a1c8f8af9cc45b5a7cbc9af9b"
	authToken    = "87f997e2aaa2f20228dd943e951aac07"
	twilioNumber = "+14849914441"
)

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func sendOTP(phoneNumber, otp string) error {
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
