package bbb_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/amirazad1/bigbluebutton-api-go/bbb"
	"github.com/amirazad1/bigbluebutton-api-go/bbb/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// -------------------- CreateMeeting --------------------
func TestCreateMeeting_Success(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/create", r.URL.Path)
		assert.Equal(t, "test123", r.URL.Query().Get("meetingID"))
		assert.Equal(t, "Test Meeting", r.URL.Query().Get("name"))
		assert.Equal(t, "ap", r.URL.Query().Get("attendeePW"))
		assert.Equal(t, "mp", r.URL.Query().Get("moderatorPW"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<response>
				<returncode>SUCCESS</returncode>
				<meetingID>test123</meetingID>
				<internalMeetingID>abc123</internalMeetingID>
			</response>`))
	})

	resp, err := client.CreateMeeting(context.Background(), &requests.CreateMeetingRequest{
		MeetingID: "test123",
		Name:      "Test Meeting",
	})

	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.Equal(t, "test123", resp.MeetingID)
	assert.Equal(t, "abc123", resp.InternalID)
}

func TestCreateMeeting_DefaultValues(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Meeting test123", r.URL.Query().Get("name"))
		assert.Equal(t, "ap", r.URL.Query().Get("attendeePW"))
		assert.Equal(t, "mp", r.URL.Query().Get("moderatorPW"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode></response>`))
	})

	_, err := client.CreateMeeting(context.Background(), &requests.CreateMeetingRequest{
		MeetingID: "test123",
	})

	require.NoError(t, err)
}

func TestCreateMeeting_ValidationErrors(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {})

	tests := []struct {
		name string
		req  *requests.CreateMeetingRequest
	}{
		{"nil request", nil},
		{"missing meetingID", &requests.CreateMeetingRequest{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.CreateMeeting(context.Background(), tt.req)
			require.Error(t, err)
		})
	}
}

// -------------------- JoinMeeting --------------------

func TestJoinMeeting_SuccessAndChecksum(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	joinURL, err := client.JoinMeeting(context.Background(), &requests.JoinMeetingRequest{
		FullName:  "Test User",
		MeetingID: "test123",
		Password:  "mp",
	})

	require.NoError(t, err)
	assert.Contains(t, joinURL, "/api/join?")
	assert.Contains(t, joinURL, "fullName=Test+User")
	assert.Contains(t, joinURL, "meetingID=test123")
	assert.Contains(t, joinURL, "password=mp")
	assert.Contains(t, joinURL, "checksum=")
}

func TestJoinMeeting_ValidationErrors(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {})

	tests := []struct {
		name string
		req  *requests.JoinMeetingRequest
	}{
		{"nil request", nil},
		{"missing meetingID", &requests.JoinMeetingRequest{Password: "mp"}},
		{"missing password", &requests.JoinMeetingRequest{MeetingID: "test123"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.JoinMeeting(context.Background(), tt.req)
			require.Error(t, err)
		})
	}
}

// -------------------- EndMeeting --------------------

func TestEndMeeting_Success(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/end", r.URL.Path)
		assert.Equal(t, "test123", r.URL.Query().Get("meetingID"))
		assert.Equal(t, "mp", r.URL.Query().Get("password"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<response>
				<returncode>SUCCESS</returncode>
				<messageKey>sentEndMeetingRequest</messageKey>
			</response>`))
	})

	resp, err := client.EndMeeting(context.Background(), &requests.EndMeetingRequest{
		MeetingID: "test123",
		Password:  "mp",
	})

	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.Equal(t, "sentEndMeetingRequest", resp.MessageKey)
}

func TestEndMeeting_ValidationErrors(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {})

	tests := []struct {
		name string
		req  *requests.EndMeetingRequest
	}{
		{"nil request", nil},
		{"missing meetingID", &requests.EndMeetingRequest{Password: "mp"}},
		{"missing password", &requests.EndMeetingRequest{MeetingID: "test123"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.EndMeeting(context.Background(), tt.req)
			require.Error(t, err)
		})
	}
}

// -------------------- GetMeetings --------------------

func TestGetMeetings_Success(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/getMeetings", r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<response>
				<returncode>SUCCESS</returncode>
				<meetings>
					<meeting>
						<meetingID>test123</meetingID>
						<meetingName>Test Meeting</meetingName>
						<running>true</running>
					</meeting>
				</meetings>
			</response>`))
	})

	resp, err := client.GetMeetings(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	require.Len(t, resp.Meetings, 1)
	assert.True(t, resp.Meetings[0].Running)
}

// -------------------- IsMeetingRunning --------------------

func TestIsMeetingRunning_Success(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/isMeetingRunning", r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<response>
				<returncode>SUCCESS</returncode>
				<running>true</running>
			</response>`))
	})

	running, err := client.IsMeetingRunning(context.Background(), "test123")
	require.NoError(t, err)
	assert.True(t, running)
}

func TestIsMeetingRunning_ValidationError(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := client.IsMeetingRunning(context.Background(), "")
	require.Error(t, err)
}

// -------------------- Server Failure --------------------

func TestServerFailure_ReturnsError(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<response>
				<returncode>FAILED</returncode>
				<message>Something went wrong</message>
			</response>`))
	})

	_, err := client.CreateMeeting(context.Background(), &requests.CreateMeetingRequest{
		MeetingID: "test123",
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "API request failed: Something went wrong")
}
