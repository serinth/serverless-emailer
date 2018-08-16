package main

import (
	"testing"
	"github.com/serinth/serverless-emailer/services"
	"github.com/serinth/serverless-emailer/api"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestSendEmailWithFallback(t *testing.T) {
	primaryEmailerMock := makeEmailerMock(false).(*services.EmailerMock)
	secondaryEmailerMock := makeEmailerMock(false).(*services.EmailerMock)
	requestStub := &api.SendEmailRequest{}

	err := sendEmailWithFallback(requestStub, primaryEmailerMock, secondaryEmailerMock)

	assert.Equal(t, 1, len(primaryEmailerMock.SendCalls()), "primary emailer should have called Send()")
	assert.Equal(t, 0, len(secondaryEmailerMock.SendCalls()), "secondary emailer should not have been called")
	assert.Nil(t, err, "sendEmailWithFallback should not have returned an error")
}

func TestSendEmailWithFallback_ShouldCallSecondaryEmailer(t *testing.T) {
	primaryEmailerMock := makeEmailerMock(true).(*services.EmailerMock)
	secondaryEmailerMock := makeEmailerMock(false).(*services.EmailerMock)
	requestStub := &api.SendEmailRequest{}

	err := sendEmailWithFallback(requestStub, primaryEmailerMock, secondaryEmailerMock)

	assert.Equal(t, 1, len(primaryEmailerMock.SendCalls()), "primary emailer should have called Send()")
	assert.Equal(t, 1, len(secondaryEmailerMock.SendCalls()), "secondary emailer should have called Send()")
	assert.Nil(t, err, "sendEmailWithFallback should not have returned an error")
}

func TestSendEmailWithFallback_ShouldReturnCriticalError_WhenBothEmailers_Fail(t *testing.T) {
	primaryEmailerMock := makeEmailerMock(true).(*services.EmailerMock)
	secondaryEmailerMock := makeEmailerMock(true).(*services.EmailerMock)
	requestStub := &api.SendEmailRequest{}

	err := sendEmailWithFallback(requestStub, primaryEmailerMock, secondaryEmailerMock)

	assert.Equal(t, 1, len(primaryEmailerMock.SendCalls()), "primary emailer should have called Send()")
	assert.Equal(t, 1, len(secondaryEmailerMock.SendCalls()), "secondary emailer should have called Send()")
	assert.Equal(t, "Critical failure with Primary and Secondary emailers", err.Error(), "should have returned critical error when both emailers failed")
}

func makeEmailerMock(shouldReturnErrorOnSend bool) services.Emailer {
	emailerMock := &services.EmailerMock{}
	emailerMock.ToFunc = func(addresses []*api.Address) services.Emailer { return emailerMock }
	emailerMock.FromFunc = func(address *api.Address) services.Emailer { return emailerMock }
	emailerMock.SubjectFunc = func(subject *string) services.Emailer { return emailerMock }
	emailerMock.ContentFunc = func(content *string) services.Emailer { return emailerMock }
	emailerMock.CCFunc = func(addresses []*api.Address) services.Emailer { return emailerMock }
	emailerMock.BCCFunc = func(addresses []*api.Address) services.Emailer { return emailerMock }

	if shouldReturnErrorOnSend {
		emailerMock.SendFunc = func() error { return errors.New("error") }
	} else {
		emailerMock.SendFunc = func() error { return nil }
	}

	return emailerMock
}
