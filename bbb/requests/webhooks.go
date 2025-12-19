/*
Package requests contains request structures for BigBlueButton API calls.
This file defines the request types for webhook-related API operations such as creating, listing, updating, and destroying webhooks.
*/

package requests

// CreateHookRequest represents the parameters for creating a webhook
type CreateHookRequest struct {
	CallbackURL string            `json:"callbackURL"`         // Required
	MeetingID   string            `json:"meetingID,omitempty"` // If empty, creates a permanent hook
	GetRaw      bool              `json:"getRaw,omitempty"`    // Return raw recording format
	Meta        map[string]string `json:"meta,omitempty"`      // Additional metadata
}

// HooksRequest represents common parameters for hook operations
type HooksRequest struct {
	HookID    string `json:"hookID,omitempty"`    // Required for all operations except create
	MeetingID string `json:"meetingID,omitempty"` // Required for meeting-specific hooks
}

// Hook represents a webhook in the system
type Hook struct {
	ID          string            `json:"id"`
	CallbackURL string            `json:"callbackURL"`
	MeetingID   string            `json:"meetingID,omitempty"` // Empty for permanent hooks
	Permanent   bool              `json:"permanent"`
	Raw         bool              `json:"raw"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   string            `json:"createdAt,omitempty"`
}
