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
		return errors
	}

	if !emailsAreValid(r) {
		errors = append(errors, "Request has invalid emails. Please check and try again.")
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
	checkEmailCollection := func(addresses []*api.Address) bool {
		if addresses != nil {
			for _, address := range addresses {
				if !isValidEmailString(address.Email) {
					return false
				}
			}
		}
		return true
	}

	return checkEmailCollection(r.CC) && checkEmailCollection(r.BCC) && checkEmailCollection(r.To) && isValidEmailString(r.From.Email)
}
