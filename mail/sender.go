package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.qq.com"
	smtpServerAddress = "smtp.qq.com:465"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type EmailConfig struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewEmailConfig(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &EmailConfig{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *EmailConfig) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	//send with ssl
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpAuthAddress,
	}
	return e.SendWithTLS(smtpServerAddress, smtpAuth, tlsConfig)

	// return e.Send(smtpServerAddress, smtpAuth)
}
