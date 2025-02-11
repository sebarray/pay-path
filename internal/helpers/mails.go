package helpers

import (
	"fmt"
	"net/smtp"
	"os"
)

const (
	SMTPHost = "smtp.gmail.com"
	SMTPPort = "587"
)

func sendMail(to, subject, body string) error {

	username := "sebarray98@gmail.com"
	password := os.Getenv("USER_PASSWORD")
	auth := smtp.PlainAuth("", username, password, SMTPHost)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	addr := fmt.Sprintf("%s:%s", SMTPHost, SMTPPort)
	err := smtp.SendMail(addr, auth, username, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

func ConfirmUserEmail(to, code string) error {
	subject := "Confirm your account"
	body := fmt.Sprintf(`Hello,<br><br>
    Please confirm your account by clicking the button below:<br><br>
    <a href="http://192.168.0.111:5173/verify?code=%s" style="display: inline-block; padding: 10px 20px; font-size: 16px; color: #ffffff; background-color: #007bff; text-decoration: none; border-radius: 5px;">Confirm Account</a><br><br>
    Best regards,`, code)
	err := sendMail(to, subject, body)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}
	fmt.Println("Email sent successfully!")
	return nil
}
