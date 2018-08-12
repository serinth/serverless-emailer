package api

const invalidEmailRequestCode = "EC_INVALID_EMAIL_REQUEST"

type errorResponse struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}

func InvalidEmailRequest(errors []string) *errorResponse {
	return &errorResponse{
		Code:   invalidEmailRequestCode,
		Errors: errors,
	}
}
