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

func TestCreateWebhook(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/hooks/create", r.URL.Path)
		assert.Equal(t, "https://example.com/callback", r.URL.Query().Get("callbackURL"))
		assert.Equal(t, "meeting_ended", r.URL.Query().Get("meetingID"))

		// Respond with success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		<response>
		  <returncode>SUCCESS</returncode>
		  <hookID>hook-123</hookID>
		  <callbackURL>https://example.com/callback</callbackURL>
		  <meetingID>meeting_ended</meetingID>
		  <permanentHook>false</permanentHook>
		</response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Test data
	req := &requests.CreateHookRequest{
		CallbackURL: "https://example.com/callback",
		MeetingID:   "meeting_ended",
	}

	// Execute
	resp, err := client.CreateHook(context.Background(), req)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.Equal(t, "hook-123", resp.HookID)
	assert.Equal(t, "https://example.com/callback", req.CallbackURL)
}

func TestListWebhooks(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/hooks/list", r.URL.Path)

		// Respond with sample webhooks data
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		<response>
		  <returncode>SUCCESS</returncode>
		  <hooks>
		    <hook>
		      <hookID>hook-123</hookID>
		      <callbackURL>https://example.com/callback</callbackURL>
		      <meetingID>meeting_ended</meetingID>
		      <permanentHook>false</permanentHook>
		    </hook>
		  </hooks>
		</response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Execute
	resp, err := client.ListHooks(context.Background())

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	require.Len(t, resp.Hooks, 1)
	assert.Equal(t, "hook-123", resp.Hooks[0].ID)
	assert.Equal(t, "https://example.com/callback", resp.Hooks[0].CallbackURL)
}

func TestDestroyWebhook(t *testing.T) {
	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "/api/hooks/destroy", r.URL.Path)
		assert.Equal(t, "hook-123", r.URL.Query().Get("hookID"))

		// Respond with success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><removed>true</removed></response>`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client, err := bbb.NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	// Test data
	hookID := "hook-123"

	// Execute
	resp, err := client.DestroyHook(context.Background(), hookID)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.True(t, resp.Removed)
}
