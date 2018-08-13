package validators

import (
	"github.com/serinth/serverless-emailer/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestStringPointer() *string {
	testString := "somestring"
	return &testString
}

func getOneValidEmail() *api.Address {
	name := "name"
	email := "me@tld.com.au"
	return &api.Address{ &name, &email }
}

func getOneInvalidEmail() *api.Address {
	invalidEmail := "adsfadf"
	return &api.Address{ nil, &invalidEmail }
}

func getValidEmailArrayPointers() []*api.Address {

	email1 := "me@gmail.com"
	email2 := "someone@blah.com.au"
	email3 := "you@derp.ca"

	return []*api.Address{
		{ nil, &email1 },
		{ nil, &email2 },
		{ nil, &email3 },
	}
}

func getInvalidEmailArrayPointers() []*api.Address {
	invalidEmail := "adsfadfasdf"

	return []*api.Address{
		{ nil, &invalidEmail },
	}
}

func TestGetSendEmailRequestErrors(t *testing.T) {
	missingMandatoryFieldsErrorMessage := "Request is missing a mandatory field. Please check the [to, from, content and subject] fields."
	invalidEmailErrors := "Request has invalid emails. Please check and try again."

	testTable := []struct {
		testName     string
		expected     []string
		request      api.SendEmailRequest
		errorMessage string
	}{
		{
			testName: "missing everything",
			expected: []string{ missingMandatoryFieldsErrorMessage },
			request:      api.SendEmailRequest{nil, nil, nil, nil, nil, nil },
			errorMessage: "Should have returned early with missing required fields error",
		},
		{
			testName: "missing required",
			expected: []string{ missingMandatoryFieldsErrorMessage },
			request:      api.SendEmailRequest{nil, getOneValidEmail(), getValidEmailArrayPointers(), getValidEmailArrayPointers(), getTestStringPointer(), getTestStringPointer() },
			errorMessage: "Should have returned required fields error",
		},
		{
			testName: "all valid fields",
			expected: nil,
			request:      api.SendEmailRequest{getValidEmailArrayPointers(), getOneValidEmail(), getValidEmailArrayPointers(), getValidEmailArrayPointers(), getTestStringPointer(), getTestStringPointer() },
			errorMessage: "Should have returned no errors",
		},
		{
			testName: "missing subject",
			expected: []string{ missingMandatoryFieldsErrorMessage },
			request:      api.SendEmailRequest{getValidEmailArrayPointers(), getOneValidEmail(), getValidEmailArrayPointers(), getValidEmailArrayPointers(), nil, getTestStringPointer() },
			errorMessage: "Should have returned required fields error",
		},
		{
			testName: "missing content",
			expected: []string{ missingMandatoryFieldsErrorMessage },
			request:      api.SendEmailRequest{getValidEmailArrayPointers(), getOneValidEmail(), getValidEmailArrayPointers(), getValidEmailArrayPointers(), getTestStringPointer(), nil },
			errorMessage: "Should have returned required fields error",
		},
		{
			testName: "invalid CC",
			expected: []string{ invalidEmailErrors },
			request:      api.SendEmailRequest{getValidEmailArrayPointers(), getOneValidEmail(), getInvalidEmailArrayPointers(), getValidEmailArrayPointers(), getTestStringPointer(), getTestStringPointer() },
			errorMessage: "Should have returned invalid emails error",
		},
		{
			testName: "invalid BCC",
			expected: []string{ invalidEmailErrors },
			request:      api.SendEmailRequest{getValidEmailArrayPointers(), getOneValidEmail(), getValidEmailArrayPointers(), getInvalidEmailArrayPointers(), getTestStringPointer(), getTestStringPointer() },
			errorMessage: "Should have returned invalid emails error",
		},
		{
			testName: "invalid To",
			expected: []string{ invalidEmailErrors },
			request:      api.SendEmailRequest{getInvalidEmailArrayPointers(), getOneValidEmail(), getValidEmailArrayPointers(), getValidEmailArrayPointers(), getTestStringPointer(), getTestStringPointer() },
			errorMessage: "Should have returned invalid emails error",
		},
		{
			testName: "invalid From",
			expected: []string{ invalidEmailErrors },
			request:      api.SendEmailRequest{getInvalidEmailArrayPointers(), getOneInvalidEmail(), getValidEmailArrayPointers(), getValidEmailArrayPointers(), getTestStringPointer(), getTestStringPointer() },
			errorMessage: "Should have returned invalid emails error",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, tt.expected, GetSendEmailRequestErrors(tt.request), tt.errorMessage)
		})
	}

}

// TODO turn into table test and test more cases
func TestIsValidEmailString(t *testing.T) {
	testString := "me@tld.com.au"

	assert.True(t, isValidEmailString(&testString) , "email should have been valid in regex match")
}


// TODO turn into table test - you get the idea
func TestEmailsAreValid(t *testing.T) {
	r := api.SendEmailRequest{getValidEmailArrayPointers(), getOneValidEmail(), getValidEmailArrayPointers(), getValidEmailArrayPointers(), nil, nil }
	assert.True(t, emailsAreValid(r), "Should have returned true for all valid emails")

	r = api.SendEmailRequest{getValidEmailArrayPointers(), getOneValidEmail(), getInvalidEmailArrayPointers(), getValidEmailArrayPointers(), nil, nil }
	assert.False(t, emailsAreValid(r))
}