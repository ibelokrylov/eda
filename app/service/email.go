package service

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"safechron/api/app/config"

	"github.com/aymerick/raymond"
)

type EmailData struct {
	Subject string
}

func SendEmail(email []string, email_data EmailData, template_name string, data map[string]string) error {
	temp, err := getTemplate(template_name)
	if err != nil {
		return err
	}

	tpl, err := raymond.Render(temp, data)
	if err != nil {
		return err
	}

	fmt.Println(tpl)

	from := config.GetEnvVariable("MAIL_USER")
	message := []byte("To: " + email[0] + "\r\n" +
		"Subject: Here is your subject\r\n" +
		"\r\n" +
		tpl + "\r\n")

	addr := config.GetEnvVariable("MAIL_SMTP") + ":" + config.GetEnvVariable("MAIL_SMTP_PORT")
	host := config.GetEnvVariable("MAIL_SMTP")

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer client.Quit()

	user := config.GetEnvVariable("MAIL_USER")
	password := config.GetEnvVariable("MAIL_PASSWORD")
	auth := smtp.PlainAuth("", user, password, host)
	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(from); err != nil {
		return err
	}
	for _, addr := range email {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(message)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully")
	return nil
}

func getTemplate(template_name string) (string, error) {
	file, err := os.ReadFile("templates/" + template_name + ".hbs")
	if err != nil {
		return "", err
	}
	return string(file), nil
}
