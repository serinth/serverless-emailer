package services

import "github.com/serinth/serverless-emailer/api"

type EmailerService struct {
	to      []*api.Address
	from    *api.Address
	cc      []*api.Address
	bcc     []*api.Address
	subject string
	content string
	isHTML  bool
	apiKey  string
	url     string
}

type Emailer interface {
	Send() error
	To(addresses []*api.Address) *EmailerService
	From(address *api.Address) *EmailerService
	CC(addresses []*api.Address) *EmailerService
	BCC(addresses []*api.Address) *EmailerService
	Subject(subject *string) *EmailerService
	Content(content *string) *EmailerService
	SetAPIKey(key string) *EmailerService
	SetUrl(url string) *EmailerService
}
