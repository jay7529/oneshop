package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

var account = Config("email.account")
var password = Config("email.password")

type mail struct {
	user   string
	passwd string
}

func New(u string, p string) mail {
	temp := mail{user: u, passwd: p}
	return temp
}

func (m mail) Send(title string, text string, toId string) {
	auth := smtp.PlainAuth("", m.user, m.passwd, "smtp.gmail.com")

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.gmail.com",
	}

	conn, err := tls.Dial("tcp", "smtp.gmail.com:465", tlsconfig)
	CheckErr(err)

	client, err := smtp.NewClient(conn, "smtp.gmail.com")
	CheckErr(err)

	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(m.user); err != nil {
		log.Panic(err)
	}

	if err = client.Rcpt(toId); err != nil {
		log.Panic(err)
	}

	w, err := client.Data()
	CheckErr(err)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", m.user, toId, title, text)

	_, err = w.Write([]byte(msg))
	CheckErr(err)

	err = w.Close()
	CheckErr(err)

	client.Quit()
}

func SendEmail(subjectLine string, content string, email string) {
	foo := New(account, password)
	foo.Send(subjectLine, content, email)
}
