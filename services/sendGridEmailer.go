package services

import (
	"bytes"
	"encoding/json"
	"github.com/serinth/serverless-emailer/api"
	"github.com/serinth/serverless-emailer/util"
	"net/http"
)

type sendGridEmailerService struct {
	to      []*api.Address
	from    *api.Address
	cc      []*api.Address
	bcc     []*api.Address
	subject string
	content string
	isHTML  bool
	apiKey  string
	url     string
	context string
}

func NewSendGridEmailer(apiKey string, url string, context string) Emailer {
	return &sendGridEmailerService{
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

func (s *sendGridEmailerService) Send() error {
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
		"application/json",
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

func (s *sendGridEmailerService) To(addresses []*api.Address) Emailer {
	s.to = addresses
	return s
}

func (s *sendGridEmailerService) From(address *api.Address) Emailer {
	s.from = address
	return s
}

func (s *sendGridEmailerService) CC(addresses []*api.Address) Emailer {
	s.cc = addresses
	return s
}

func (s *sendGridEmailerService) BCC(addresses []*api.Address) Emailer {
	s.bcc = addresses
	return s
}

func (s *sendGridEmailerService) Subject(subject *string) Emailer {
	s.subject = *subject
	return s
}

func (s *sendGridEmailerService) Content(content *string) Emailer {
	s.content = *content
	return s
}

func (s *sendGridEmailerService) SetAPIKey(key string) Emailer {
	s.apiKey = key
	return s
}

func (s *sendGridEmailerService) SetUrl(url string) Emailer {
	s.url = url
	return s
}
