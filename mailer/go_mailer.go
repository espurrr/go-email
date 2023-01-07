package mailer

import (
	"bytes"
	"goemailclient/config"
	"goemailclient/models"
	"html/template"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func setHTMLBody(mail models.Mail) string {
	//set html template
	t := template.New("index.html")

	var err error
	//parse html file you want
	t, err = t.ParseFiles("index.html")
	if err != nil {
		log.Error(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, mail); err != nil {
		log.Error(err)
	}
	// return tpl as a string
	htmlBody := tpl.String()
	return htmlBody
}

func SendMail(mail models.Mail) {

	// get environment variable
	godotenv.Load()

	// set html page
	htmlBody := setHTMLBody(mail)

	// set email
	dialer, msg := config.ConfigEmail(mail, htmlBody)

	// send E-Mail
	if err := dialer.DialAndSend(msg); err != nil {
		log.Panic(err) // calls panic() after logging
	} else {
		log.Info("Email Sent Successfully!")
	}
}
