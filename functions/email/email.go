package main

import (
	"errors"
	"fmt"
	"github.com/serinth/serverless-emailer/api"
	"github.com/serinth/serverless-emailer/services"
	"github.com/serinth/serverless-emailer/util"
	log "github.com/sirupsen/logrus"
)

func local() {

	fmt.Println("Running Locally...")

	cfg = util.LoadConfig()

	to := "serinth+test1@gmail.com"
	from := "postmaster@sandboxfd35f37be9664c8abfc2f2cdb66a6961.mailgun.org"
	subject := "test"
	content := "test content"

	req := &api.SendEmailRequest{
		To: []*api.Address{
			{nil, &to},
		},
		From:    &api.Address{Name: nil, Email: &from},
		Subject: &subject,
		Content: &content,
		CC:      nil,
		BCC:     nil,
	}

	emailUsingSendGrid(req, cfg)

	emailUsingMailgun(req, cfg)

}

func sendEmail(req *api.SendEmailRequest, cfg *util.Config) error {
	if err := emailUsingSendGrid(req, cfg); err != nil {
		log.Errorf("Email with SendGrid Failed with error: %v", err)
		fallbackErr := emailUsingMailgun(req, cfg)
		if fallbackErr != nil {
			log.Errorf("Fallback with Mailgun failed with error: %v", fallbackErr)
			return errors.New("Critical failure with SendGrid and Mailgun")
		}
	}

	return nil
}

func emailUsingSendGrid(req *api.SendEmailRequest, cfg *util.Config) error {
	emailer := services.NewSendGridEmailer(cfg.SendGridAPIKey, cfg.SendGridURL, cfg.MetricsCommandName)
	err := emailer.
		To(req.To).
		From(req.From).
		Subject(req.Subject).
		Content(req.Content).
		Send()

	return err
}

func emailUsingMailgun(req *api.SendEmailRequest, cfg *util.Config) error {
	emailer := services.NewMailgunEmailer(cfg.MailGunAPIKey, cfg.MailGunURL, cfg.MetricsCommandName)
	err := emailer.
		To(req.To).
		From(req.From).
		Subject(req.Subject).
		Content(req.Content).
		Send()

	return err
}
