package bbb

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// NewTestClient creates a new test client with the provided handler.
func NewTestClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()

	ts := newTestServer(t, handler)
	client, err := NewClient(ts.URL, "test-secret")
	require.NoError(t, err)

	return client
}

// newTestServer creates a new test server with the provided handler.
func newTestServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()

	ts := httptest.NewServer(handler)
	t.Cleanup(func() {
		ts.Close()
	})

	return ts
}
