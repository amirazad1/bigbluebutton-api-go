/*
Package responses contains response structures for BigBlueButton API calls.
This file defines the response types for webhook-related API operations such as create, list, destroy, and update webhooks.
*/

package responses

// CreateHookResponse represents the response from creating a webhook
type CreateHookResponse struct {
	BaseResponseImpl
	HookID string `xml:"hookID"`
}

// HooksResponse represents a list of webhooks
type HooksResponse struct {
	BaseResponseImpl
	Hooks []HookDetails `xml:"hooks>hook"`
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
	BaseResponseImpl
	Removed bool `xml:"removed"`
}
