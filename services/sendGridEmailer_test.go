package services

import (
	"github.com/serinth/serverless-emailer/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapApiAddressToSendGridAddresses(t *testing.T) {
	email := "email@email.com"
	name := "name"
	expected := []*sendGridAddress{
		{Email: email},
		{Name: name, Email: email},
	}

	addresses := []*api.Address{
		{nil, &email},
		{&name, &email},
	}

	assert.Equal(t, expected, mapApiAddressToSendGridAddresses(addresses), "Mapping to Mailgun addresses did not produce the proper form string")
}

// TODO Integration Tests here for Send(), point to mock services on dev environment
