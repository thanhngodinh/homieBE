package send_email

import (
	"fmt"
	"net/smtp"
)

func SendPasswordResetEmail(email string, newPassword string) error {
	subject := "Password Reset\r\n"
	body := fmt.Sprintf("Your new password is: %s", newPassword)

	return sendEmail(email, subject, body)
}

func SendVerifyEmail(email string, code string) error {
	subject := "Verify Email\r\n"
	body := fmt.Sprintf("Your verify email is: %s", code)

	return sendEmail(email, subject, body)
}

func sendEmail(email string, subject string, body string) error {
	from := "thanh.ngodinh2000@gmail.com"
	password := "yulc kzeg rery eewk"
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	message := []byte(subject + "\r\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpServer)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{email}, message)
	if err != nil {
		return err
	}

	return nil
}
