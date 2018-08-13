package services

import (
	"bytes"
	"fmt"
	"github.com/serinth/serverless-emailer/util"
	"net/http"
	"github.com/serinth/serverless-emailer/api"
)

func NewSendGridEmailer(apiKey string, url string) Emailer {
	return &EmailerService{
		apiKey: apiKey,
		url:    url,
		isHTML: false,
	}
}

func (s *EmailerService) Send() error {
	// TODO serialize this instead of building a string, placeholder
	data := []byte(fmt.Sprintf(`{"personalizations": [{"to": [{"email": "example@example.com"}]}],"from": {"email": "example@example.com"},"subject": "Hello, World!","content": [{"type": "text/plain", "value": "Heya!"}]}`))
	err := util.HystrixPost(
		http.MethodPost,
		s.url,
		bytes.NewBuffer(data),
		util.AuthCredentials{IsBasicAuth: false, APIKey: s.apiKey},
		"SendGridRequest",
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *EmailerService) To(addresses []*api.Address) *EmailerService {
	s.to = addresses
	return s
}

func (s *EmailerService) From(address *api.Address) *EmailerService {
	s.from = address
	return s
}

func (s *EmailerService) CC(addresses []*api.Address) *EmailerService {
	s.cc = addresses
	return s
}

func (s *EmailerService) BCC(addresses []*api.Address) *EmailerService {
	s.bcc = addresses
	return s
}

func (s *EmailerService) Subject(subject *string) *EmailerService {
	s.subject = *subject
	return s
}

func (s *EmailerService) Content(content *string) *EmailerService {
	s.content = *content
	return s
}

func (s *EmailerService) SetAPIKey(key string) *EmailerService {
	s.apiKey = key
	return s
}

func (s *EmailerService) SetUrl(url string) *EmailerService {
	s.url = url
	return s
}
