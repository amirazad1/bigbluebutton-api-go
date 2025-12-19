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

func TestGetRecordings(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/getRecordings", r.URL.Path)
		assert.Equal(t, "test123", r.URL.Query().Get("meetingID"))

		// Respond with sample recordings data
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
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Test data
	req := &requests.GetRecordingsRequest{
		MeetingID: "test123",
	}

	// Execute
	resp, err := client.GetRecordings(context.Background(), req)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	require.Len(t, resp.Recordings, 1)
	assert.Equal(t, "recording-123", resp.Recordings[0].RecordID)
	assert.Equal(t, "test123", resp.Recordings[0].MeetingID)
	assert.True(t, resp.Recordings[0].Published)
}

func TestPublishRecordings(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/publishRecordings", r.URL.Path)
		assert.Equal(t, "recording-123", r.URL.Query().Get("recordID"))
		assert.Equal(t, "true", r.URL.Query().Get("publish"))

		// Respond with success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><published>true</published></response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Test data
	req := &requests.PublishRecordingsRequest{
		RecordID: "recording-123",
		Publish:  true,
	}

	// Execute
	resp, err := client.PublishRecordings(context.Background(), req)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.True(t, resp.Published)
}

func TestDeleteRecordings(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/deleteRecordings", r.URL.Path)
		assert.Equal(t, "recording-123", r.URL.Query().Get("recordID"))

		// Respond with success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><deleted>true</deleted></response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Execute
	recordID := "recording-123"
	resp, err := client.DeleteRecordings(context.Background(), recordID)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.True(t, resp.Deleted)
}

func TestUpdateRecordings(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/updateRecordings", r.URL.Path)
		assert.Equal(t, "recording-123", r.URL.Query().Get("recordID"))
		assert.Equal(t, "Updated Recording", r.URL.Query().Get("meta_name"))

		// Respond with success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><updated>true</updated></response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Test data
	req := &requests.UpdateRecordingsRequest{
		RecordID: "recording-123",
		Meta: map[string]string{
			"name": "Updated Recording",
		},
	}

	// Execute
	resp, err := client.UpdateRecordings(context.Background(), req)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.True(t, resp.Updated)
}
