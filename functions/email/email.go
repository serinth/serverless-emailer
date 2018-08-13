package main

import (
	//"encoding/json"
	"fmt"
	"github.com/serinth/serverless-emailer/api"
	"github.com/serinth/serverless-emailer/services"
	"github.com/serinth/serverless-emailer/util"
)

func local() {

	fmt.Println("Running Locally...")

	cfg = util.LoadConfig()

	to := "serinth+test1@gmail.com"
	from := "serinth@gmail.com"
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
	emailer := services.NewSendGridEmailer(cfg.SendGridAPIKey, cfg.SendGridURL, cfg.MetricsCommandName)
	err := emailer.
		To(req.To).
		From(req.From).
		Subject(req.Subject).
		Content(req.Content).
		Send()

	fmt.Println("Error if any: %v", err)
}

func sendEmail(r api.SendEmailRequest) error {
	return nil
}
