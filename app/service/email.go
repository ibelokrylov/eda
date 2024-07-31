package service

import (
	"eda/app/config"
	"net/smtp"
	"os"

	"github.com/aymerick/raymond"
)

type EmailData struct {
	Subject string
}

func renderTemplate(filePath string, context map[string]string) (string, error) {
	templateData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	template := string(templateData)
	result, err := raymond.Render(template, context)
	if err != nil {
		return "", err
	}
	return result, nil
}

func SendEmail(to []string, subject, filePath string, context map[string]string) error {
	from := config.GetEnvVariable("MAIL_USER")
	password := config.GetEnvVariable("MAIL_PASSWORD")

	smtpHost := config.GetEnvVariable("MAIL_SMTP")
	smtpPort := config.GetEnvVariable("MAIL_SMTP_PORT")

	body, err := renderTemplate(filePath, context)
	if err != nil {
		return err
	}
	message := []byte(
		"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
			body)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	println(auth)

	error := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	return error
}

func SendEmailCodeConfirmRegistration(to string, code string) error {
	context := map[string]string{
		"title": "Code registration",
		"code":  code,
	}
	return SendEmail([]string{to}, "Code confirm registration", "templates/registration.hbs", context)
}
