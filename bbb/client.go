/*
Package bbb provides a Go client for the BigBlueButton API.

This file contains the core Client struct and methods for initializing the client,
configuring HTTP options, generating checksums, and making HTTP requests.
*/

package bbb

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client represents a BigBlueButton API client.
type Client struct {
	baseURL    string
	secret     string
	httpClient *http.Client
}

// Option configures a Client.
type Option func(*Client) error

// NewClient creates a new BigBlueButton API client.
func NewClient(baseURL, secret string, options ...Option) (*Client, error) {
	// Ensure baseURL ends with a slash
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	if !strings.HasSuffix(baseURL, "api/") {
		baseURL += "api/"
	}

	// Create default client
	c := &Client{
		baseURL: baseURL,
		secret:  secret,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	// Apply options
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, fmt.Errorf("applying option: %w", err)
		}
	}

	return c, nil
}

// WithHTTPClient sets the HTTP client for the BigBlueButton client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

// WithTimeout sets the timeout for the HTTP client.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		c.httpClient.Timeout = timeout
		return nil
	}
}

// generateChecksum generates a checksum for the given API call and query parameters.
func (c *Client) generateChecksum(apiCall string, params url.Values) string {
	queryString := params.Encode()
	sha := sha1.New()
	sha.Write([]byte(apiCall + queryString + c.secret))
	return hex.EncodeToString(sha.Sum(nil))
}

// doRequest performs an HTTP request to the BigBlueButton API.
func (c *Client) doRequest(ctx context.Context, action string, params url.Values, result interface{}) error {
	// Build the URL with the correct API path
	u := fmt.Sprintf("%s%s", c.baseURL, action)

	// Add checksum to parameters
	checksum := c.generateChecksum(action, params)
	params.Set("checksum", checksum)

	// Build the full URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", u, params.Encode())

	// Print the URL for debugging (remove in production)
	fmt.Printf("Making request to: %s\n", fullURL)

	// Create the request
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Make the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body for error details
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	// Print the response status and body for debugging
	fmt.Printf("Response Status: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(body))

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	// Parse the XML response
	if err := xml.Unmarshal(body, result); err != nil {
		return fmt.Errorf("parsing response: %w, response body: %s", err, string(body))
	}

	return nil
}
