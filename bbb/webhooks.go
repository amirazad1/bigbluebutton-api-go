/*
Package bbb provides functionality for managing BigBlueButton webhooks.
This file contains methods for creating, listing, updating, and destroying webhooks.
*/

package bbb

import (
	"context"
	"net/url"

	"github.com/amirazad1/bigbluebutton-api-go/bbb/requests"
	"github.com/amirazad1/bigbluebutton-api-go/bbb/responses"
)

// CreateHook creates a new webhook
// If MeetingID is empty, creates a permanent hook that receives all events
func (c *Client) CreateHook(ctx context.Context, req *requests.CreateHookRequest) (*responses.CreateHookResponse, error) {
	if req == nil {
		return nil, NewError(ErrInvalidParam, "request cannot be nil")
	}

	if req.CallbackURL == "" {
		return nil, NewError(ErrMissingParam, "callbackURL is required")
	}

	params := url.Values{
		"callbackURL": {req.CallbackURL},
	}

	// Add optional parameters
	if req.MeetingID != "" {
		params.Set("meetingID", req.MeetingID)
	}
	if req.GetRaw {
		params.Set("getRaw", "true")
	}

	// Add metadata
	for k, v := range req.Meta {
		params.Set("meta_"+k, v)
	}

	var response responses.CreateHookResponse
	if err := c.doRequest(ctx, "hooks/create", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListHooks retrieves all webhooks
func (c *Client) ListHooks(ctx context.Context) (*responses.HooksResponse, error) {
	params := url.Values{}
	var response responses.HooksResponse

	if err := c.doRequest(ctx, "hooks/list", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ListHooksForMeeting retrieves all webhooks for a specific meeting
func (c *Client) ListHooksForMeeting(ctx context.Context, meetingID string) (*responses.HooksResponse, error) {
	if meetingID == "" {
		return nil, NewError(ErrMissingParam, "meetingID is required")
	}

	params := url.Values{
		"meetingID": {meetingID},
	}

	var response responses.HooksResponse
	if err := c.doRequest(ctx, "hooks/list", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DestroyHook removes a webhook
func (c *Client) DestroyHook(ctx context.Context, hookID string) (*responses.DestroyHookResponse, error) {
	if hookID == "" {
		return nil, NewError(ErrMissingParam, "hookID is required")
	}

	params := url.Values{
		"hookID": {hookID},
	}

	var response responses.DestroyHookResponse
	if err := c.doRequest(ctx, "hooks/destroy", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateHook updates an existing webhook
func (c *Client) UpdateHook(ctx context.Context, hookID string, updates map[string]string) (*responses.CreateHookResponse, error) {
	if hookID == "" {
		return nil, NewError(ErrMissingParam, "hookID is required")
	}

	params := url.Values{
		"hookID": {hookID},
	}

	// Add updates
	for k, v := range updates {
		params.Set(k, v)
	}

	var response responses.CreateHookResponse
	if err := c.doRequest(ctx, "hooks/update", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
