// Package bbb provides a Go client for the BigBlueButton API.
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

// APIResponse represents a basic API response from BigBlueButton
type APIResponse struct {
	ReturnCode string `xml:"returncode"`
	Version    string `xml:"version"`
	Message    string `xml:"message,omitempty"`
	MessageKey string `xml:"messageKey,omitempty"`
}

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

	// Ensure the URL ends with 'api/'
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

// GetAPIVersion returns the version of the BigBlueButton server.
func (c *Client) GetAPIVersion(ctx context.Context) (string, error) {
	var response APIResponse

	if err := c.doRequest(ctx, "api", url.Values{}, &response); err != nil {
		return "", fmt.Errorf("failed to get API version: %w", err)
	}

	if response.ReturnCode != "SUCCESS" {
		return "", fmt.Errorf("API error: %s", response.Message)
	}

	return response.Version, nil
}

// doRequest performs an HTTP request to the BigBlueButton API.
func (c *Client) doRequest(ctx context.Context, action string, params url.Values, result interface{}) error {
    // Build the URL with the correct API path
    u := fmt.Sprintf("%sapi/%s", c.baseURL, action)
    
    // Add checksum to parameters
    checksum := c.generateChecksum(action, params)
    params.Set("checksum", checksum)
    
    // Build the full URL with query parameters
    fullURL := fmt.Sprintf("%s?%s", u, params.Encode())
    
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
    
    // Check status code
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }
    
    // Read and parse the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("reading response body: %w", err)
    }
    
    // Parse the XML response
    if err := xml.Unmarshal(body, result); err != nil {
        return fmt.Errorf("parsing response: %w", err)
    }
    
    return nil
}
