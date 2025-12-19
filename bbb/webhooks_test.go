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

// -------------------- CreateWebhook --------------------
func TestCreateWebhook(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/hooks/create", r.URL.Path)
		assert.Equal(t, "https://example.com/callback", r.URL.Query().Get("callbackURL"))
		assert.Equal(t, "meeting_ended", r.URL.Query().Get("meetingID"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		<response>
		  <returncode>SUCCESS</returncode>
		  <hookID>hook-123</hookID>
		  <callbackURL>https://example.com/callback</callbackURL>
		  <meetingID>meeting_ended</meetingID>
		  <permanentHook>false</permanentHook>
		</response>`))
	})

	req := &requests.CreateHookRequest{
		CallbackURL: "https://example.com/callback",
		MeetingID:   "meeting_ended",
	}

	resp, err := client.CreateHook(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.Equal(t, "hook-123", resp.HookID)
	assert.Equal(t, "https://example.com/callback", req.CallbackURL)
}

// -------------------- ListWebhooks --------------------
func TestListWebhooks(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/hooks/list", r.URL.Path)

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
	})

	resp, err := client.ListHooks(context.Background())

	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	require.Len(t, resp.Hooks, 1)
	assert.Equal(t, "hook-123", resp.Hooks[0].ID)
	assert.Equal(t, "https://example.com/callback", resp.Hooks[0].CallbackURL)
}

// -------------------- DestroyWebhook --------------------
func TestDestroyWebhook(t *testing.T) {
	client := bbb.NewTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/hooks/destroy", r.URL.Path)
		assert.Equal(t, "hook-123", r.URL.Query().Get("hookID"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response><returncode>SUCCESS</returncode><removed>true</removed></response>`))
	})

	hookID := "hook-123"

	resp, err := client.DestroyHook(context.Background(), hookID)

	require.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.ReturnCode)
	assert.True(t, resp.Removed)
}
