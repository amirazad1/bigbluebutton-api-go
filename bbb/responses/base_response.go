package responses

// BaseResponse defines the common interface for all API responses
type BaseResponse interface {
	GetReturnCode() string
	GetMessage() string
}

// BaseResponseImpl provides a base implementation of BaseResponse
type BaseResponseImpl struct {
	ReturnCode string `xml:"returncode"`
	MessageKey string `xml:"messageKey,omitempty"`
	Message    string `xml:"message,omitempty"`
}

// GetReturnCode returns the return code from the response
func (r *BaseResponseImpl) GetReturnCode() string {
	return r.ReturnCode
}

// GetMessage returns the message from the response
func (r *BaseResponseImpl) GetMessage() string {
	if r.Message != "" {
		return r.Message
	}
	return r.MessageKey
}
