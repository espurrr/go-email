package models

//model mail
type Mail struct {
	Address string `json:"address"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

//model mails
type Mails []Mail
