package validators

import (
	"github.com/serinth/serverless-emailer/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSendEmailRequestErrors(t *testing.T) {
	testTable := []struct {
		testName     string
		expected     []string
		request      api.SendEmailRequest
		errorMessage string
	}{
		{
			testName: "missing everything",
			expected: []string{
				"Request is missing a mandatory field. Please check the [to, from, content and subject] fields.",
				"Request has invalid emails. Please check and try again.",
				"Subject cannot be empty.",
				"Content cannot be empty.",
			},
			request:      api.SendEmailRequest{nil, nil, nil, nil, nil, nil},
			errorMessage: "Should have returned array with all the possible errors",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expected, GetSendEmailRequestErrors(tt.request), tt.errorMessage)
		})
	}

}
