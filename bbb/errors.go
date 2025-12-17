package bbb

import "fmt"

// Error represents a BigBlueButton API error.
type Error struct {
	Code    string
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("BigBlueButton error: %s (code: %s)", e.Message, e.Code)
}

// Common error codes and messages.
const (
	// General errors
	ErrInvalidURL          = "invalid_url"
	ErrInvalidResponse     = "invalid_response"
	ErrRequestFailed       = "request_failed"
	ErrUnauthorized        = "unauthorized"
	ErrNotFound            = "not_found"
	ErrInternalServerError = "internal_server_error"

	// API specific errors
	ErrChecksumMismatch = "checksum_mismatch"
	ErrMissingParam     = "missing_parameter"
	ErrInvalidParam     = "invalid_parameter"
)

// NewError creates a new Error with the given code and message.
func NewError(code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// IsError checks if the error is a BigBlueButton error with the given code.
func IsError(err error, code string) bool {
	if err == nil {
		return false
	}

	if bbbErr, ok := err.(*Error); ok {
		return bbbErr.Code == code
	}

	return false
}
