package config

import (
	"crypto/tls"
	"goemailclient/models"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	gomail "gopkg.in/mail.v2"
)

func ConfigEmail(mail models.Mail, htmlBody string) (*gomail.Dialer, *gomail.Message) {

	//get Email configs of sender from env variables
	senders_email := os.Getenv("MAIL_USERNAME")
	mail_from_name := os.Getenv("MAIL_FROM_NAME")
	password := os.Getenv("MAIL_PASSWORD")
	mailer := os.Getenv("MAIL_MAILER")
	mailer_port := os.Getenv("MAIL_PORT")

	mailer_port_int, err := strconv.Atoi(mailer_port)
	if err == nil {
		// port is of type int
	} else {
		log.Error("Mailer port could not be converted to int")
	}

	//create New Message Function
	msg := gomail.NewMessage()

	// Set E-Mail sender
	msg.SetHeader("From", msg.FormatAddress(senders_email, mail_from_name))

	// Set E-Mail receivers
	msg.SetHeader("To", mail.Address)

	// Set E-Mail subject
	msg.SetHeader("Subject", mail.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	msg.SetBody("text/html", htmlBody)

	// Set E-Mail Cc Bcc
	// msg.SetAddressHeader("Cc", "abc@example.com", "abc")

	// Set attachments
	// msg.Attach("cat.jpg")

	// Settings for SMTP server
	dialer := gomail.NewDialer(mailer, mailer_port_int, senders_email, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	//return msg and dialer
	return dialer, msg
}
