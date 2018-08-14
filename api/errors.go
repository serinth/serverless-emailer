package api

const invalidEmailRequestCode = "EC_INVALID_EMAIL_REQUEST"
const internalErrorCode = "EC_INTERNAL_ERROR"

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

func InternalError(err string) *errorResponse {
	return &errorResponse{
		Code:   internalErrorCode,
		Errors: []string{err},
	}
}
