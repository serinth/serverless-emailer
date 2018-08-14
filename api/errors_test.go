package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidEmailRequest(t *testing.T) {
	errorData := []string{"error1", "error2"}
	expected := &errorResponse{
		Code:   invalidEmailRequestCode,
		Errors: errorData,
	}

	assert.Equal(t, expected, InvalidEmailRequest(errorData), "Error response returned was not as expected")
}

func TestInternalError(t *testing.T) {
	errorData := "error"
	expected := &errorResponse{
		Code:   internalErrorCode,
		Errors: []string{errorData},
	}

	assert.Equal(t, expected, InternalError(errorData), "Error response returned was not as expected")
}
