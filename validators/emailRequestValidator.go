package validators

import (
	"github.com/serinth/serverless-emailer/api"
	"regexp"
)

// returns an array of reasons why request was not valid
func GetSendEmailRequestErrors(r api.SendEmailRequest) []string {
	var errors []string

	if isMissingRequiredFields(r) {
		errors = append(errors, "Request is missing a mandatory field. Please check the [to, from, content and subject] fields.")
	}

	if !emailsAreValid(r) {
		errors = append(errors, "Request has invalid emails. Please check and try again.")
	}

	if r.Subject == nil || len(*r.Subject) == 0 {
		errors = append(errors, "Subject cannot be empty.")
	}

	if r.Content == nil || len(*r.Content) == 0 {
		errors = append(errors, "Content cannot be empty.")
	}

	return errors
}

func isMissingRequiredFields(r api.SendEmailRequest) bool {
	return r.To == nil || r.From == nil || r.Content == nil || r.Subject == nil
}

// chose to use simple regex pattern matching for emails. See ARCHITECTURE.md for the reasoning.
func isValidEmailString(address *string) bool {
	if address == nil {
		return false
	}

	match, _ := regexp.MatchString(".+\\@.+\\..+", *address)
	return match
}

func emailsAreValid(r api.SendEmailRequest) bool {
	checkEmailCollection := func(emails []*string) bool {
		if emails != nil {
			for _, email := range emails {
				if !isValidEmailString(email) {
					return false
				}
			}
		}
		return true
	}

	if !isValidEmailString(r.To) || !isValidEmailString(r.From) {
		return false
	}

	return checkEmailCollection(r.CC) && checkEmailCollection(r.BCC)
}
