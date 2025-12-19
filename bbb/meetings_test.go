package bbb_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amirazad1/bigbluebutton-api-go/bbb"
	"github.com/amirazad1/bigbluebutton-api-go/bbb/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateMeeting(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/create", r.URL.Path)
		assert.Equal(t, "test123", r.URL.Query().Get("meetingID"))
		assert.Equal(t, "Test Meeting", r.URL.Query().Get("name"))
		assert.Equal(t, "ap", r.URL.Query().Get("attendeePW"))
		assert.Equal(t, "mp", r.URL.Query().Get("moderatorPW"))

		// Respond with success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><meetingID>test123</meetingID><internalMeetingID>abc123</internalMeetingID><createTime>1234567890</createTime><voiceBridge>12345</voiceBridge></response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Test data
	req := &requests.CreateMeetingRequest{
		MeetingID:    "test123",
		Name:         "Test Meeting",
		AttendeePW:   "ap",
		ModeratorPW:  "mp",
	}

	// Execute
	resp, err := client.CreateMeeting(context.Background(), req)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.Equal(t, "test123", resp.MeetingID)
	assert.Equal(t, "abc123", resp.InternalID)
}

func TestJoinMeeting(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Just respond with OK, the actual join URL is constructed by the client
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Test data
	req := &requests.JoinMeetingRequest{
		FullName:    "Test User",
		MeetingID:   "test123",
		Password:    "mp",
	}

	// Execute
	joinURL, err := client.JoinMeeting(context.Background(), req)

	// Verify
	require.NoError(t, err)
	assert.Contains(t, joinURL, "/api/join?")
	assert.Contains(t, joinURL, "fullName=Test+User")
	assert.Contains(t, joinURL, "meetingID=test123")
	assert.Contains(t, joinURL, "password=mp")
}

func TestEndMeeting(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/end", r.URL.Path)
		assert.Equal(t, "test123", r.URL.Query().Get("meetingID"))
		assert.Equal(t, "mp", r.URL.Query().Get("password"))

		// Respond with success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><messageKey>sentEndMeetingRequest</messageKey><message>End meeting request sent</message></response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Test data
	req := &requests.EndMeetingRequest{
		MeetingID: "test123",
		Password:  "mp",
	}

	// Execute
	resp, err := client.EndMeeting(context.Background(), req)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.Equal(t, "sentEndMeetingRequest", resp.MessageKey)
}

func TestGetMeetings(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/getMeetings", r.URL.Path)

		// Respond with sample meetings data
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		<response>
		  <returncode>SUCCESS</returncode>
		  <meetings>
		    <meeting>
		      <meetingID>test123</meetingID>
		      <meetingName>Test Meeting</meetingName>
		      <createTime>1234567890</createTime>
		      <attendeePW>ap</attendeePW>
		      <moderatorPW>mp</moderatorPW>
		      <hasBeenForciblyEnded>false</hasBeenForciblyEnded>
		      <running>true</running>
		      <participantCount>1</participantCount>
		    </meeting>
		  </meetings>
		</response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Execute
	resp, err := client.GetMeetings(context.Background())

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	require.Len(t, resp.Meetings, 1)
	assert.Equal(t, "test123", resp.Meetings[0].MeetingID)
	assert.Equal(t, "Test Meeting", resp.Meetings[0].MeetingName)
	assert.True(t, resp.Meetings[0].Running)
}

func TestIsMeetingRunning(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/isMeetingRunning", r.URL.Path)
		assert.Equal(t, "test123", r.URL.Query().Get("meetingID"))

		// Respond with meeting running status
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><running>true</running></response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Execute
	running, err := client.IsMeetingRunning(context.Background(), "test123")

	// Verify
	require.NoError(t, err)
	assert.True(t, running)
}
