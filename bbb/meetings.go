package bbb

import (
	"context"
	"net/url"
	"strconv"

	"github.com/amirazad1/bigbluebutton-api-go/bbb/requests"
	"github.com/amirazad1/bigbluebutton-api-go/bbb/responses"
)

// CreateMeeting creates a new meeting with the specified parameters.
func (c *Client) CreateMeeting(ctx context.Context, req *requests.CreateMeetingRequest) (*responses.CreateMeetingResponse, error) {
	if req == nil {
		return nil, NewError(ErrInvalidParam, "request cannot be nil")
	}

	// Set default values if not provided
	if req.MeetingID == "" {
		return nil, NewError(ErrMissingParam, "meetingID is required")
	}
	if req.Name == "" {
		req.Name = "Meeting " + req.MeetingID
	}
	if req.AttendeePW == "" {
		req.AttendeePW = "ap"
	}
	if req.ModeratorPW == "" {
		req.ModeratorPW = "mp"
	}

	// Prepare parameters
	params := url.Values{}
	params.Set("name", req.Name)
	params.Set("meetingID", req.MeetingID)
	params.Set("attendeePW", req.AttendeePW)
	params.Set("moderatorPW", req.ModeratorPW)

	// Add optional parameters if provided
	if req.Welcome != "" {
		params.Set("welcome", req.Welcome)
	}
	if req.DialNumber != "" {
		params.Set("dialNumber", req.DialNumber)
	}
	if req.VoiceBridge != "" {
		params.Set("voiceBridge", req.VoiceBridge)
	}
	if req.WebVoice != "" {
		params.Set("webVoice", req.WebVoice)
	}
	if req.LogoutURL != "" {
		params.Set("logoutURL", req.LogoutURL)
	}
	if req.MaxParticipants > 0 {
		params.Set("maxParticipants", strconv.Itoa(req.MaxParticipants))
	}

	// Add boolean flags
	params.Set("record", boolToStr(req.Record))
	params.Set("autoStartRecording", boolToStr(req.AutoStartRecording))
	params.Set("allowStartStopRecording", boolToStr(req.AllowStartStopRecording))
	params.Set("webcamsOnlyForModerator", boolToStr(req.WebcamsOnlyForModerator))
	params.Set("muteOnStart", boolToStr(req.MuteOnStart))

	// Add lock settings
	params.Set("lockSettingsDisableCam", boolToStr(req.LockSettingsDisableCam))
	params.Set("lockSettingsDisableMic", boolToStr(req.LockSettingsDisableMic))
	params.Set("lockSettingsDisablePrivateChat", boolToStr(req.LockSettingsDisablePrivateChat))
	params.Set("lockSettingsDisablePublicChat", boolToStr(req.LockSettingsDisablePublicChat))
	params.Set("lockSettingsLockedLayout", boolToStr(req.LockSettingsLockedLayout))
	params.Set("lockSettingsLockOnJoin", boolToStr(req.LockSettingsLockOnJoin))
	params.Set("lockSettingsLockOnJoinConfigurable", boolToStr(req.LockSettingsLockOnJoinConfigurable))

	// Add metadata
	for k, v := range req.Meta {
		params.Set("meta_"+k, v)
	}

	// Make the API call
	var response responses.CreateMeetingResponse
	if err := c.doRequest(ctx, "create", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// JoinMeeting generates a join URL for a meeting.
func (c *Client) JoinMeeting(ctx context.Context, req *requests.JoinMeetingRequest) (string, error) {
	if req == nil {
		return "", NewError(ErrInvalidParam, "request cannot be nil")
	}

	// Validate required parameters
	if req.MeetingID == "" {
		return "", NewError(ErrMissingParam, "meetingID is required")
	}
	if req.Password == "" {
		return "", NewError(ErrMissingParam, "password is required")
	}
	if req.FullName == "" {
		req.FullName = "User"
	}

	// Prepare parameters
	params := url.Values{}
	params.Set("meetingID", req.MeetingID)
	params.Set("password", req.Password)
	params.Set("fullName", req.FullName)

	// Add optional parameters
	if req.UserID != "" {
		params.Set("userID", req.UserID)
	}
	if req.CreateTime != "" {
		params.Set("createTime", req.CreateTime)
	}

	// Add user data
	for k, v := range req.UserData {
		params.Set("userdata_"+k, v)
	}

	// Generate the join URL
	checksum := c.generateChecksum("join", params)
	params.Set("checksum", checksum)

	// Build the full URL
	joinURL := c.baseURL + "join?" + params.Encode()

	return joinURL, nil
}

// EndMeeting ends a running meeting.
func (c *Client) EndMeeting(ctx context.Context, req *requests.EndMeetingRequest) (*responses.EndMeetingResponse, error) {
	if req == nil {
		return nil, NewError(ErrInvalidParam, "request cannot be nil")
	}

	// Validate required parameters
	if req.MeetingID == "" {
		return nil, NewError(ErrMissingParam, "meetingID is required")
	}
	if req.Password == "" {
		return nil, NewError(ErrMissingParam, "password is required")
	}

	// Prepare parameters
	params := url.Values{}
	params.Set("meetingID", req.MeetingID)
	params.Set("password", req.Password)

	// Make the API call
	var response responses.EndMeetingResponse
	if err := c.doRequest(ctx, "end", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetMeetingInfo retrieves information about a specific meeting.
func (c *Client) GetMeetingInfo(ctx context.Context, meetingID, password string) (*responses.GetMeetingInfoResponse, error) {
	if meetingID == "" {
		return nil, NewError(ErrMissingParam, "meetingID is required")
	}
	if password == "" {
		return nil, NewError(ErrMissingParam, "password is required")
	}

	// Prepare parameters
	params := url.Values{}
	params.Set("meetingID", meetingID)
	params.Set("password", password)

	// Make the API call
	var response responses.GetMeetingInfoResponse
	if err := c.doRequest(ctx, "getMeetingInfo", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetMeetings retrieves a list of all running meetings.
func (c *Client) GetMeetings(ctx context.Context) (*responses.GetMeetingsResponse, error) {
	// No parameters needed for getMeetings
	params := url.Values{}

	// Make the API call
	var response responses.GetMeetingsResponse
	if err := c.doRequest(ctx, "getMeetings", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// IsMeetingRunning checks if a meeting is running.
func (c *Client) IsMeetingRunning(ctx context.Context, meetingID string) (bool, error) {
	if meetingID == "" {
		return false, NewError(ErrMissingParam, "meetingID is required")
	}

	// Prepare parameters
	params := url.Values{}
	params.Set("meetingID", meetingID)

	// Response structure for isMeetingRunning
	type isMeetingRunningResponse struct {
		ReturnCode string `xml:"returncode"`
		Running    bool   `xml:"running"`
	}

	// Make the API call
	var response isMeetingRunningResponse
	if err := c.doRequest(ctx, "isMeetingRunning", params, &response); err != nil {
		return false, err
	}

	return response.Running, nil
}

// boolToStr converts a boolean to "true" or "false" string.
func boolToStr(b bool) string {
	return strconv.FormatBool(b)
}
