/*
Package responses contains response structures for BigBlueButton API calls.
This file defines the response types for meeting-related API operations such as create, join, end, get info, and list meetings.
*/

package responses

// CreateMeetingResponse represents the response from the create meeting API
type CreateMeetingResponse struct {
	BaseResponseImpl
	MeetingID     string `xml:"meetingID"`
	InternalID    string `xml:"internalMeetingID"`
	ParentID      string `xml:"parentMeetingID"`
	AttendeePW    string `xml:"attendeePW"`
	ModeratorPW   string `xml:"moderatorPW"`
	CreateTime    string `xml:"createTime"`
	VoiceBridge   string `xml:"voiceBridge"`
	DialNumber    string `xml:"dialNumber"`
	CreateDate    string `xml:"createDate"`
	HasUserJoined bool   `xml:"hasUserJoined"`
	Duration      int    `xml:"duration"`
}

// JoinMeetingResponse represents the response from the join meeting API
type JoinMeetingResponse struct {
	BaseResponseImpl
	UserID       string `xml:"user_id"`
	MeetingID    string `xml:"meeting_id"`
	AuthToken    string `xml:"auth_token"`
	SessionToken string `xml:"session_token"`
	GuestStatus  string `xml:"guestStatus"`
	URL          string `xml:"url"`
}

// EndMeetingResponse represents the response from the end meeting API
type EndMeetingResponse struct {
	BaseResponseImpl
}

// GetMeetingInfoResponse represents the response from the get meeting info API
type GetMeetingInfoResponse struct {
	BaseResponseImpl
	MeetingName           string            `xml:"meetingName"`
	MeetingID             string            `xml:"meetingID"`
	InternalID            string            `xml:"internalMeetingID"`
	CreateTime            string            `xml:"createTime"`
	CreateDate            string            `xml:"createDate"`
	VoiceBridge           string            `xml:"voiceBridge"`
	DialNumber            string            `xml:"dialNumber"`
	AttendeePW            string            `xml:"attendeePW"`
	ModeratorPW           string            `xml:"moderatorPW"`
	Running               bool              `xml:"running"`
	Recording             bool              `xml:"recording"`
	HasBeenForciblyEnded  bool              `xml:"hasBeenForciblyEnded"`
	StartTime             string            `xml:"startTime"`
	EndTime               string            `xml:"endTime"`
	ParticipantCount      int               `xml:"participantCount"`
	ListenerCount         int               `xml:"listenerCount"`
	VoiceParticipantCount int               `xml:"voiceParticipantCount"`
	VideoCount            int               `xml:"videoCount"`
	Duration              int               `xml:"duration"`
	HasUserJoined         bool              `xml:"hasUserJoined"`
	Metadata              map[string]string `xml:"metadata"`
	ModeratorCount        int               `xml:"moderatorCount"`
}

// Meeting represents a meeting in the getMeetings response
type Meeting struct {
	MeetingID             string            `xml:"meetingID"`
	MeetingName           string            `xml:"meetingName"`
	CreateTime            int64             `xml:"createTime"`
	VoiceBridge           string            `xml:"voiceBridge"`
	DialNumber            string            `xml:"dialNumber"`
	AttendeePW            string            `xml:"attendeePW"`
	ModeratorPW           string            `xml:"moderatorPW"`
	HasUserJoined         bool              `xml:"hasUserJoined"`
	HasBeenForciblyEnded  bool              `xml:"hasBeenForciblyEnded"`
	Running               bool              `xml:"running"`
	ParticipantCount      int               `xml:"participantCount"`
	ListenerCount         int               `xml:"listenerCount"`
	VoiceParticipantCount int               `xml:"voiceParticipantCount"`
	VideoCount            int               `xml:"videoCount"`
	Duration              int               `xml:"duration"`
	CreateDate            string            `xml:"createDate"`
	StartTime             string            `xml:"startTime"`
	EndTime               string            `xml:"endTime"`
	Metadata              map[string]string `xml:"metadata"`
}

// GetMeetingsResponse represents the response from the getMeetings API
type GetMeetingsResponse struct {
	BaseResponseImpl
	Meetings []Meeting `xml:"meetings>meeting"`
}
