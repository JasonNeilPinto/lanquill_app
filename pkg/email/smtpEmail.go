package email

import (
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/smtp"
	"os"
)

var LANQUILL_SUPPORT_EMAIL_PASSWORD = os.Getenv("LANQUILL_SUPPORT_EMAIL_PASSWORD")
var LANQUILL_SUPPORT_EMAIL = os.Getenv("LANQUILL_SUPPORT_EMAIL")

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

func SendEmail(Subject string, to []string, Body string) error {
	// Sender data.
	from := LANQUILL_SUPPORT_EMAIL
	password := LANQUILL_SUPPORT_EMAIL_PASSWORD

	// smtp server configuration.
	smtpHost := "smtp.office365.com"
	smtpPort := "587"

	// Authentication.
	conn, err := net.Dial("tcp", smtpHost+":"+smtpPort)
	if err != nil {
		log.Println("Error: ", err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Println("Error: ", err)
	}

	tlsconfig := &tls.Config{
		ServerName: smtpHost,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		log.Println("Error: ", err)
	}
	auth := LoginAuth(from, password)

	if err = c.Auth(auth); err != nil {
		log.Println("Error: ", err)
	}

	// Message.
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte("Subject:" + Subject + "\n" + mime + Body)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	return nil
}
