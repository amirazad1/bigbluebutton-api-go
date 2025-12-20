# BigBlueButton API Client for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/amirazad1/bigbluebutton-api-go.svg)](https://pkg.go.dev/github.com/amirazad1/bigbluebutton-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/amirazad1/bigbluebutton-api-go)](https://goreportcard.com/report/github.com/amirazad1/bigbluebutton-api-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/amirazad1/bigbluebutton-api-go)](https://golang.org/)

A comprehensive, type-safe Go client library for interacting with the BigBlueButton API, designed for reliability, performance, and ease of use.

## ✨ Features

- **Full API Coverage**: Complete implementation of BigBlueButton API v2.4+
- **Thread-Safe**: Safe for concurrent use by multiple goroutines
- **Context Support**: Built-in support for timeouts and request cancellation
- **Robust Error Handling**: Comprehensive error types and messages
- **Well-Documented**: Extensive GoDoc with practical examples
- **High Test Coverage**: Thoroughly tested with 100% code coverage
- **Modular Design**: Clean separation of concerns with `requests` and `responses` packages
- **Webhook Support**: Full support for BigBlueButton webhooks
- **No External Dependencies**: Lightweight and dependency-free

## 📦 Installation

```bash
go get github.com/amirazad1/bigbluebutton-api-go
```

## 🚀 Quick Start

### Prerequisites
- Go 1.16 or higher
- A running BigBlueButton server (v2.4+)
- API secret from your BigBlueButton server

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/amirazad1/bigbluebutton-api-go/bbb"
	"github.com/amirazad1/bigbluebutton-api-go/bbb/requests"
)

func main() {
	// Initialize client with custom timeout
	client, err := bbb.NewClient(
		"https://your-bbb-server/bigbluebutton", // Your BigBlueButton server URL
		"your-api-secret-here",                   // Your API secret
		bbb.WithTimeout(15*time.Second),          // Optional: Set custom timeout
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a new meeting
	meeting, err := client.CreateMeeting(context.Background(), &requests.CreateMeetingRequest{
		Name:            "Team Planning Session",
		MeetingID:       "unique-meeting-id-123",
		AttendeePW:      "attendee-pass",
		ModeratorPW:     "moderator-pass",
		Welcome:         "Welcome to our planning meeting!",
		Record:          true,
		MaxParticipants: 25,
		Meta: map[string]string{
			"meeting-purpose": "sprint-planning",
			"department":     "engineering",
		},
	})

	if err != nil {
		log.Fatalf("Failed to create meeting: %v", err)
	}

	// Generate join URL for moderator
	joinURL, err := client.JoinMeeting(context.Background(), &requests.JoinMeetingRequest{
		MeetingID: meeting.MeetingID,
		Password:  "moderator-pass",
		FullName:  "Jane Doe",
		UserID:    "user-456",
		UserData: map[string]string{
			"role":     "moderator",
			"presence": "present",
		},
	})

	if err != nil {
		log.Fatalf("Failed to generate join URL: %v", err)
	}

	fmt.Printf("Meeting created successfully!\n")
	fmt.Printf("Meeting ID: %s\n", meeting.MeetingID)
	fmt.Printf("Join URL: %s\n", joinURL)
}
```

## 🔧 Configuration

### Client Options

The client supports various configuration options:

```go
import (
	"net/http"
	"time"

	"github.com/amirazad1/bigbluebutton-api-go/bbb"
)

// 1. With custom HTTP client
customClient := &http.Client{
	Timeout: 20 * time.Second,
	// Add custom transport, timeouts, etc.
}

// 2. With custom timeout
client, _ := bbb.NewClient(
	"https://your-bbb-server/bigbluebutton",
	"your-api-secret",
	bbb.WithTimeout(30*time.Second),      // Set request timeout
)

// 3. With custom HTTP client
client, _ := bbb.NewClient(
	"https://your-bbb-server/bigbluebutton",
	"your-api-secret",
	bbb.WithHTTPClient(customClient),     // Use custom HTTP client
)
```

## 📚 API Coverage

### Meetings
- [x] Create meeting
- [x] Join meeting
- [x] End meeting
- [x] Get meeting info
- [x] List all meetings
- [x] Check if meeting is running

### Recordings
- [x] Get recordings
- [x] Publish/unpublish recordings
- [x] Delete recordings
- [x] Update recording metadata

### Webhooks
- [x] Create hook
- [x] List all hooks
- [x] List hooks for meeting
- [x] Update hook
- [x] Destroy hook

## Usage Examples

### Create a Meeting

```go
meeting, err := client.CreateMeeting(context.Background(), &requests.CreateMeetingRequest{
    Name:           "Team Meeting",
    MeetingID:      "meeting-123",
    AttendeePW:     "ap",
    ModeratorPW:    "mp",
    Welcome:        "Welcome to our team meeting!",
    Record:         true,
    MaxParticipants: 50,
    Meta: map[string]string{
        "meeting-purpose": "weekly-sync",
    },
})
if err != nil {
    log.Fatalf("Failed to create meeting: %v", err)
}
```

### Join a Meeting

```go
joinURL, err := client.JoinMeeting(context.Background(), &requests.JoinMeetingRequest{
    MeetingID: "meeting-123",
    Password:  "mp", // Moderator password (use "ap" for attendees)
    FullName:  "John Doe",
    UserID:    "user-123",
    UserData: map[string]string{
        "role": "moderator",
    },
})
if err != nil {
    log.Fatalf("Failed to generate join URL: %v", err)
}
// Redirect user to joinURL
```

### Manage Recordings

```go
// Get all recordings
recordings, err := client.GetRecordings(context.Background(), &requests.GetRecordingsRequest{
    MeetingID: "meeting-123",
    State:     "any",
})

// Publish a recording
_, err = client.PublishRecordings(context.Background(), &requests.PublishRecordingsRequest{
    RecordID: "recording-123",
    Publish:  true,
})

// Delete a recording
_, err = client.DeleteRecordings(context.Background(), "recording-123")
```

### Webhooks

```go
// Create a webhook
hook, err := client.CreateHook(context.Background(), &requests.CreateHookRequest{
    CallbackURL: "https://your-server.com/webhook",
    MeetingID:   "meeting-123", // Omit for global webhook
    GetRaw:      true,
    Meta: map[string]string{
        "description": "Webhook for meeting events",
    },
})

// List all webhooks
hooks, err := client.ListHooks(context.Background())

// Remove a webhook
_, err = client.DestroyHook(context.Background(), "hook-123")
```

## Webhook Payload Example

When an event occurs, your webhook URL will receive a POST request with a JSON payload like:

```json
{
  "header": {
    "name": "meeting-created",
    "timestamp": "2023-01-01T00:00:00Z"
  },
  "payload": {
    "meeting": {
      "id": "meeting-123",
      "name": "Test Meeting",
      "recorded": true
    }
  }
}
```

## Documentation

Full API documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/amirazad1/bigbluebutton-api-go).

## Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
