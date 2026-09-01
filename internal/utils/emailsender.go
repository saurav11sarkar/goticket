package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/saurav11sarkar/ticket/internal/config"
	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	cfg *config.Config
}

func NewEmailSender(cfg *config.Config) *EmailSender {
	return &EmailSender{cfg: cfg}
}

func (es *EmailSender) SendEmail(email, subject, html string) error {
	if es.cfg.EmailHost == "" || es.cfg.EmailPass == "" || es.cfg.EmailAddress == "" {
		return errors.New("email host, email password and email address must be set")
	}

	port, err := strconv.Atoi(es.cfg.EmailPort)
	if err != nil || port < 1 || port > 65535 {
		return errors.New("email port must be a number between 1 and 65535")
	}

	message := gomail.NewMessage()
	from := es.cfg.EmailFrom
	if from == "" {
		from = es.cfg.EmailAddress
	}

	message.SetHeader("From", from)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", html)

	dialer := gomail.NewDialer(es.cfg.EmailHost, port, es.cfg.EmailAddress, es.cfg.EmailPass)
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("email send failed: %w", err)
	}
	return nil
}
