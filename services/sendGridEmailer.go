package services

import (
	"bytes"
	"encoding/json"
	"github.com/serinth/serverless-emailer/api"
	"github.com/serinth/serverless-emailer/util"
	"net/http"
)

func NewSendGridEmailer(apiKey string, url string, context string) Emailer {
	return &EmailerService{
		apiKey:  apiKey,
		url:     url,
		isHTML:  false,
		context: context,
	}
}

// https://sendgrid.com/docs/API_Reference/api_v3.html

type sendGridAddress struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

type sendGridContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type personalizations struct {
	To  []*sendGridAddress `json:"to"`
	CC  []*sendGridAddress `json:"cc,omitempty"`
	BCC []*sendGridAddress `json:"bcc,omitempty"`
}

type sendGridRequest struct {
	Personalizations []*personalizations `json:"personalizations"`
	Subject          string              `json:"subject"`
	From             *sendGridAddress    `json:"from"`
	Content          []*sendGridContent  `json:"content"`
}

func (s *EmailerService) Send() error {
	from := &sendGridAddress{Email: *s.from.Email}
	if s.from.Name != nil {
		from.Name = *s.from.Name
	}

	req := &sendGridRequest{
		Personalizations: []*personalizations{
			{
				To:  mapApiAddressToSendGridAddresses(s.to),
				CC:  mapApiAddressToSendGridAddresses(s.cc),
				BCC: mapApiAddressToSendGridAddresses(s.bcc),
			},
		},
		From:    from,
		Subject: s.subject,
		Content: []*sendGridContent{
			{Type: "text/plain", Value: s.content},
		},
	}

	data, _ := json.Marshal(req)

	err := util.HystrixPost(
		http.MethodPost,
		s.url,
		bytes.NewBuffer(data),
		util.AuthCredentials{IsBasicAuth: false, APIKey: s.apiKey},
		s.context,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func mapApiAddressToSendGridAddresses(addresses []*api.Address) []*sendGridAddress {
	var sendGridAddresses []*sendGridAddress
	for _, a := range addresses {
		sendGridAddress := &sendGridAddress{Email: *a.Email}

		if a.Name != nil {
			sendGridAddress.Name = *a.Name
		}

		sendGridAddresses = append(sendGridAddresses, sendGridAddress)
	}

	return sendGridAddresses
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
