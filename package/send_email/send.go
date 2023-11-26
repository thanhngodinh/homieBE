package send_email

import (
	"fmt"
	"net/smtp"
)

func SendPasswordResetEmail(email string, newPassword string) error {
	subject := "Password Reset"
	body := fmt.Sprintf("Your new password is: %s", newPassword)

	return sendEmail(email, subject, body)
}

func SendVerifyEmail(email string, code string) error {
	subject := "Verify Email"
	body := fmt.Sprintf("Your verify email is: %s", code)

	return sendEmail(email, subject, body)
}

func sendEmail(email string, subject string, body string) error {
	from := "thanh.ngodinh2000@gmail.com"
	password := "yulc kzeg rery eewk"
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	auth := smtp.PlainAuth("", from, password, smtpServer)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{email}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
