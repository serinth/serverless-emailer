package services

import "github.com/serinth/serverless-emailer/api"

//go:generate moq -out emailer_mock.go . Emailer

type Emailer interface {
	Send() error
	To(addresses []*api.Address) Emailer
	From(address *api.Address) Emailer
	CC(addresses []*api.Address) Emailer
	BCC(addresses []*api.Address) Emailer
	Subject(subject *string) Emailer
	Content(content *string) Emailer
	SetAPIKey(key string) Emailer
	SetUrl(url string) Emailer
}
