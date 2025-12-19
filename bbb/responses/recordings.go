/*
Package responses contains response structures for BigBlueButton API calls.
This file defines the response types for recordings-related API operations such as get, publish, delete, and update recordings.
*/

package responses

// RecordingFormat represents a recording format in the API response
type RecordingFormat struct {
	Type       string            `xml:"type,attr"`
	URL        string            `xml:"url,omitempty"`
	Length     int               `xml:"length,omitempty"`
	Processing string            `xml:"processing,omitempty"`
	Preview    *RecordingPreview `xml:"preview,omitempty"`
}

// RecordingPreview represents the preview images for a recording
type RecordingPreview struct {
	Images []RecordingImage `xml:"images>image"`
}

// RecordingImage represents an image in the recording preview
type RecordingImage struct {
	Alt    string `xml:"alt,attr"`
	Height int    `xml:"height,attr"`
	Width  int    `xml:"width,attr"`
	Link   string `xml:"link,attr"`
}

// Recording represents a recording in the API response
type Recording struct {
	RecordID        string             `xml:"recordID"`
	MeetingID       string             `xml:"meetingID"`
	Name            string             `xml:"name"`
	Published       bool               `xml:"published"`
	State           string             `xml:"state"`
	StartTime       string             `xml:"startTime"`
	EndTime         string             `xml:"endTime"`
	Participants    int                `xml:"participants"`
	Metadata        map[string]string  `xml:"metadata"`
	Playback        *RecordingPlayback `xml:"playback"`
	RawSize         int64              `xml:"rawSize"`
	Size            int64              `xml:"size"`
	IsBreakout      bool               `xml:"isBreakout"`
	ParentMeetingID string             `xml:"parentMeetingID"`
}

// RecordingPlayback represents the playback information for a recording
type RecordingPlayback struct {
	Type    string            `xml:"type,attr"`
	URL     string            `xml:"url"`
	Length  int               `xml:"length,omitempty"`
	Preview *RecordingPreview `xml:"preview,omitempty"`
}

// GetRecordingsResponse represents the response from the getRecordings API
type GetRecordingsResponse struct {
	BaseResponseImpl
	Recordings []Recording `xml:"recordings>recording"`
}

// PublishRecordingsResponse represents the response from the publishRecordings API
type PublishRecordingsResponse struct {
	BaseResponseImpl
	Published bool `xml:"published"`
}

// DeleteRecordingsResponse represents the response from the deleteRecordings API
type DeleteRecordingsResponse struct {
	BaseResponseImpl
	Deleted bool `xml:"deleted"`
}

// UpdateRecordingsResponse represents the response from the updateRecordings API
type UpdateRecordingsResponse struct {
	BaseResponseImpl
	Updated bool `xml:"updated"`
}
