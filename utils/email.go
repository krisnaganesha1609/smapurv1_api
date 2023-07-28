package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"
	m "smapurv1_api/models"
	s "smapurv1_api/setup"

	"github.com/k3a/html2text"

	gomail "gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func TemplateParser(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *m.Users, data *EmailData, emailTemp string) {
	config, err := s.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	from := config.EmailFrom
	smtpPass := config.SMTPPass
	smtpUser := config.SMTPUser
	to := user.Email
	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort

	var body bytes.Buffer

	template, err := TemplateParser("templates")
	if err != nil {
		log.Fatal("Could not parse templates", err)
	}

	template.ExecuteTemplate(&body, emailTemp, &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email", err)
	}
}
