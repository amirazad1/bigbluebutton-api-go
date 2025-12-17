package requests

// GetRecordingsRequest represents the parameters for getting recordings
type GetRecordingsRequest struct {
	MeetingID string `json:"meetingID,omitempty"`
	RecordID  string `json:"recordID,omitempty"`
	State     string `json:"state,omitempty"` // "any", "processing", "processed", "published", "unpublished", "deleted"
	Meta      string `json:"meta,omitempty"`  // A meta parameter to filter by
	Offset    int    `json:"offset,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

// PublishRecordingsRequest represents the parameters for publishing/unpublishing recordings
type PublishRecordingsRequest struct {
	RecordID string `json:"recordID"`
	Publish  bool   `json:"publish"`
}

// DeleteRecordingsRequest represents the parameters for deleting recordings
type DeleteRecordingsRequest struct {
	RecordID string `json:"recordID"`
}

// UpdateRecordingsRequest represents the parameters for updating recording metadata
type UpdateRecordingsRequest struct {
	RecordID string            `json:"recordID"`
	Meta     map[string]string `json:"meta,omitempty"`
}
