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

// -------------------- GetRecordings --------------------
func TestGetRecordings_Success(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/getRecordings", r.URL.Path)
		assert.Equal(t, "test123", r.URL.Query().Get("meetingID"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<response>
				<returncode>SUCCESS</returncode>
				<recordings>
					<recording>
						<recordID>recording-123</recordID>
						<meetingID>test123</meetingID>
						<name>Test Recording</name>
						<published>true</published>
						<state>published</state>
						<startTime>1234567890</startTime>
						<endTime>1234568990</endTime>
						<playback>
							<format>
								<type>presentation</type>
								<url>https://example.com/playback/presentation/1.0/playback.html?meetingId=test123</url>
								<length>1100</length>
							</format>
						</playback>
					</recording>
				</recordings>
			</response>`))
	})

	req := &requests.GetRecordingsRequest{MeetingID: "test123"}
	resp, err := client.GetRecordings(context.Background(), req)
	require.NoError(t, err)

	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	require.Len(t, resp.Recordings, 1)
	assert.Equal(t, "recording-123", resp.Recordings[0].RecordID)
	assert.Equal(t, "test123", resp.Recordings[0].MeetingID)
	assert.True(t, resp.Recordings[0].Published)
}

// -------------------- PublishRecordings --------------------

func TestPublishRecordings_Success(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/publishRecordings", r.URL.Path)
		assert.Equal(t, "recording-123", r.URL.Query().Get("recordID"))
		assert.Equal(t, "true", r.URL.Query().Get("publish"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><published>true</published></response>`))
	})

	req := &requests.PublishRecordingsRequest{
		RecordID: "recording-123",
		Publish:  true,
	}
	resp, err := client.PublishRecordings(context.Background(), req)
	require.NoError(t, err)

	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.True(t, resp.Published)
}

// -------------------- DeleteRecordings --------------------

func TestDeleteRecordings_Success(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/deleteRecordings", r.URL.Path)
		assert.Equal(t, "recording-123", r.URL.Query().Get("recordID"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><deleted>true</deleted></response>`))
	})

	resp, err := client.DeleteRecordings(context.Background(), "recording-123")
	require.NoError(t, err)

	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.True(t, resp.Deleted)
}

// -------------------- UpdateRecordings --------------------

func TestUpdateRecordings_Success(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/updateRecordings", r.URL.Path)
		assert.Equal(t, "recording-123", r.URL.Query().Get("recordID"))
		assert.Equal(t, "Updated Recording", r.URL.Query().Get("meta_name"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><updated>true</updated></response>`))
	})

	req := &requests.UpdateRecordingsRequest{
		RecordID: "recording-123",
		Meta: map[string]string{
			"name": "Updated Recording",
		},
	}
	resp, err := client.UpdateRecordings(context.Background(), req)
	require.NoError(t, err)

	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.True(t, resp.Updated)
}
