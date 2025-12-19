/*
Package requests contains request structures for BigBlueButton API calls.
This file defines the request types for meeting-related API operations such as creating, joining, and ending meetings.
*/

package requests

// CreateMeetingRequest represents the parameters for creating a new meeting
type CreateMeetingRequest struct {
	Name                               string            `json:"name"`
	MeetingID                          string            `json:"meetingID"`
	AttendeePW                         string            `json:"attendeePW"`
	ModeratorPW                        string            `json:"moderatorPW"`
	Welcome                            string            `json:"welcome,omitempty"`
	DialNumber                         string            `json:"dialNumber,omitempty"`
	VoiceBridge                        string            `json:"voiceBridge,omitempty"`
	WebVoice                           string            `json:"webVoice,omitempty"`
	LogoutURL                          string            `json:"logoutURL,omitempty"`
	MaxParticipants                    int               `json:"maxParticipants,omitempty"`
	Record                             bool              `json:"record,omitempty"`
	AutoStartRecording                 bool              `json:"autoStartRecording,omitempty"`
	AllowStartStopRecording            bool              `json:"allowStartStopRecording,omitempty"`
	WebcamsOnlyForModerator            bool              `json:"webcamsOnlyForModerator,omitempty"`
	MuteOnStart                        bool              `json:"muteOnStart,omitempty"`
	LockSettingsDisableCam             bool              `json:"lockSettingsDisableCam,omitempty"`
	LockSettingsDisableMic             bool              `json:"lockSettingsDisableMic,omitempty"`
	LockSettingsDisablePrivateChat     bool              `json:"lockSettingsDisablePrivateChat,omitempty"`
	LockSettingsDisablePublicChat      bool              `json:"lockSettingsDisablePublicChat,omitempty"`
	LockSettingsLockedLayout           bool              `json:"lockSettingsLockedLayout,omitempty"`
	LockSettingsLockOnJoin             bool              `json:"lockSettingsLockOnJoin,omitempty"`
	LockSettingsLockOnJoinConfigurable bool              `json:"lockSettingsLockOnJoinConfigurable,omitempty"`
	Meta                               map[string]string `json:"meta,omitempty"`
}

// JoinMeetingRequest represents the parameters for joining a meeting
type JoinMeetingRequest struct {
	FullName   string            `json:"fullName"`
	MeetingID  string            `json:"meetingID"`
	Password   string            `json:"password"`
	UserID     string            `json:"userId,omitempty"`
	CreateTime string            `json:"createTime,omitempty"`
	UserData   map[string]string `json:"userData,omitempty"`
}

// EndMeetingRequest represents the parameters for ending a meeting
type EndMeetingRequest struct {
	MeetingID string `json:"meetingID"`
	Password  string `json:"password"`
}
