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

	// test code locally

}

func sendEmail(req *api.SendEmailRequest, cfg *util.Config) error {
	sendGridEmailer := services.NewSendGridEmailer(cfg.SendGridAPIKey, cfg.SendGridURL, cfg.MetricsCommandName)
	mailgunEmailer := services.NewMailgunEmailer(cfg.MailGunAPIKey, cfg.MailGunURL, cfg.MetricsCommandName)

	return sendEmailWithFallback(req, sendGridEmailer, mailgunEmailer)
}

func sendEmailWithFallback(req *api.SendEmailRequest, primaryEmailService services.Emailer, secondaryEmailService services.Emailer) error {
	if err := sendEmailFromRequest(primaryEmailService, req); err != nil {
		log.Errorf("Primary Emailer Failed with error: %v", err)
		fallbackErr := sendEmailFromRequest(secondaryEmailService, req)
		if fallbackErr != nil {
			log.Errorf("Fallback Emailer failed with error: %v", fallbackErr)
			return errors.New("Critical failure with Primary and Secondary emailers")
		}
	}

	return nil
}

func sendEmailFromRequest(emailer services.Emailer, req *api.SendEmailRequest) error {
	err := emailer.
		To(req.To).
		From(req.From).
		Subject(req.Subject).
		Content(req.Content).
		CC(req.CC).
		BCC(req.BCC).
		Send()

	return err
}
