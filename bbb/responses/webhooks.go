package responses

// CreateHookResponse represents the response from creating a webhook
type CreateHookResponse struct {
	ReturnCode string `xml:"returncode"`
	HookID     string `xml:"hookID"`
	MessageKey string `xml:"messageKey,omitempty"`
	Message    string `xml:"message,omitempty"`
}

// HooksResponse represents a list of webhooks
type HooksResponse struct {
	ReturnCode string        `xml:"returncode"`
	Hooks      []HookDetails `xml:"hooks>hook"`
	MessageKey string        `xml:"messageKey,omitempty"`
	Message    string        `xml:"message,omitempty"`
}

// HookDetails represents detailed information about a webhook
type HookDetails struct {
	ID          string            `xml:"hookID"`
	CallbackURL string            `xml:"callbackURL"`
	MeetingID   string            `xml:"meetingID,omitempty"`
	Permanent   bool              `xml:"permanent"`
	Raw         bool              `xml:"raw"`
	Metadata    map[string]string `xml:"metadata>entry"` // This might need custom unmarshaling
	CreatedAt   string            `xml:"createdAt,omitempty"`
}

// DestroyHookResponse represents the response from destroying a webhook
type DestroyHookResponse struct {
	ReturnCode string `xml:"returncode"`
	Removed    bool   `xml:"removed"`
	MessageKey string `xml:"messageKey,omitempty"`
	Message    string `xml:"message,omitempty"`
}
