/*
Package bbb provides functionality for managing BigBlueButton recordings.
This file contains methods for retrieving, publishing, deleting, and updating recordings.
*/

package bbb

import (
	"context"
	"net/url"
	"strconv"

	"github.com/amirazad1/bigbluebutton-api-go/bbb/requests"
	"github.com/amirazad1/bigbluebutton-api-go/bbb/responses"
)

// GetRecordings retrieves recordings that are available for playback.
func (c *Client) GetRecordings(ctx context.Context, req *requests.GetRecordingsRequest) (*responses.GetRecordingsResponse, error) {
	if req == nil {
		req = &requests.GetRecordingsRequest{}
	}

	params := url.Values{}

	// Add optional parameters
	if req.MeetingID != "" {
		params.Set("meetingID", req.MeetingID)
	}
	if req.RecordID != "" {
		params.Set("recordID", req.RecordID)
	}
	if req.State != "" {
		params.Set("state", req.State)
	}
	if req.Meta != "" {
		params.Set("meta", req.Meta)
	}
	if req.Offset > 0 {
		params.Set("offset", strconv.Itoa(req.Offset))
	}
	if req.Limit > 0 {
		params.Set("limit", strconv.Itoa(req.Limit))
	}

	// Make the API call
	var response responses.GetRecordingsResponse
	if err := c.doRequest(ctx, "getRecordings", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PublishRecordings publishes or unpublishes a recording.
func (c *Client) PublishRecordings(ctx context.Context, req *requests.PublishRecordingsRequest) (*responses.PublishRecordingsResponse, error) {
	if req == nil {
		return nil, NewError(ErrInvalidParam, "request cannot be nil")
	}

	if req.RecordID == "" {
		return nil, NewError(ErrMissingParam, "recordID is required")
	}

	params := url.Values{
		"recordID": {req.RecordID},
		"publish":  {strconv.FormatBool(req.Publish)},
	}

	// Make the API call
	var response responses.PublishRecordingsResponse
	if err := c.doRequest(ctx, "publishRecordings", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteRecordings deletes a recording.
func (c *Client) DeleteRecordings(ctx context.Context, recordID string) (*responses.DeleteRecordingsResponse, error) {
	if recordID == "" {
		return nil, NewError(ErrMissingParam, "recordID is required")
	}

	params := url.Values{
		"recordID": {recordID},
	}

	// Make the API call
	var response responses.DeleteRecordingsResponse
	if err := c.doRequest(ctx, "deleteRecordings", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateRecordings updates a recording's metadata.
func (c *Client) UpdateRecordings(ctx context.Context, req *requests.UpdateRecordingsRequest) (*responses.UpdateRecordingsResponse, error) {
	if req == nil {
		return nil, NewError(ErrInvalidParam, "request cannot be nil")
	}

	if req.RecordID == "" {
		return nil, NewError(ErrMissingParam, "recordID is required")
	}

	params := url.Values{
		"recordID": {req.RecordID},
	}

	// Add metadata
	for k, v := range req.Meta {
		params.Set("meta_"+k, v)
	}

	// Make the API call
	var response responses.UpdateRecordingsResponse
	if err := c.doRequest(ctx, "updateRecordings", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
