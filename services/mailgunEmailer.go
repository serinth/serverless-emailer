package services

import (
	"fmt"
	"github.com/serinth/serverless-emailer/api"
	"github.com/serinth/serverless-emailer/util"
	"net/http"
	"strings"
	"net/url"
)

type mailgunEmailerService struct {
	to      []*api.Address
	from    *api.Address
	cc      []*api.Address
	bcc     []*api.Address
	subject string
	text    string
	apiKey  string
	url     string
	context string
}

func NewMailgunEmailer(apiKey string, url string, context string) Emailer {
	return &mailgunEmailerService{
		apiKey:  apiKey,
		url:     url,
		context: context,
	}
}

func (s *mailgunEmailerService) Send() error {
	var from string

	if s.from.Name != nil {
		from = fmt.Sprintf("%s <%s>", *s.from.Name, *s.from.Email)
	} else {
		from = *s.from.Email
	}


	data := url.Values{}
	data.Set("from", from)
	data.Set("to", mapApiAddressToMailgunAddresses(s.to))

	if s.cc != nil {
		data.Set("cc", mapApiAddressToMailgunAddresses(s.cc))
	}

	if s.bcc != nil {
		data.Set("bcc", mapApiAddressToMailgunAddresses(s.bcc))
	}

	data.Set("subject", s.subject)
	data.Set("text", s.text)

	err := util.HystrixPost(
		http.MethodPost,
		s.url,
		strings.NewReader(data.Encode()),
		util.AuthCredentials{IsBasicAuth: true, User: "api", Password: s.apiKey}, //mailgun always uses "api" as the user
		s.context,
		"application/x-www-form-urlencoded",
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func mapApiAddressToMailgunAddresses(addresses []*api.Address) string {
	var formatted []string
	for _, a := range addresses {
		if a.Name != nil {
			formatted = append(formatted, fmt.Sprintf("%s <%s>", *a.Name, *a.Email))
		} else {
			formatted = append(formatted, *a.Email)
		}
	}

	return strings.Join(formatted, ",")
}

func (s *mailgunEmailerService) To(addresses []*api.Address) Emailer {
	s.to = addresses
	return s
}

func (s *mailgunEmailerService) From(address *api.Address) Emailer {
	s.from = address
	return s
}

func (s *mailgunEmailerService) CC(addresses []*api.Address) Emailer {
	s.cc = addresses
	return s
}

func (s *mailgunEmailerService) BCC(addresses []*api.Address) Emailer {
	s.bcc = addresses
	return s
}

func (s *mailgunEmailerService) Subject(subject *string) Emailer {
	s.subject = *subject
	return s
}

func (s *mailgunEmailerService) Content(content *string) Emailer {
	s.text = *content
	return s
}

func (s *mailgunEmailerService) SetAPIKey(key string) Emailer {
	s.apiKey = key
	return s
}

func (s *mailgunEmailerService) SetUrl(url string) Emailer {
	s.url = url
	return s
}
