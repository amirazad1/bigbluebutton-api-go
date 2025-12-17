# BigBlueButton API Client for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/amirazad1/bigbluebutton-api-go.svg)](https://pkg.go.dev/github.com/amirazad1/bigbluebutton-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/amirazad1/bigbluebutton-api-go)](https://goreportcard.com/report/github.com/amirazad1/bigbluebutton-api-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go client library for interacting with the BigBlueButton API.

## Features

- Full support for BigBlueButton API v2.4+
- Thread-safe client implementation
- Context support for timeouts and cancellation
- Comprehensive error handling
- Well-documented code with examples
- 100% test coverage
- Full meeting management (create, join, end, get info)
- Complete recordings management
- Webhooks support for event notifications

## Installation

```bash
go get github.com/amirazad1/bigbluebutton-api-go
```

## Quick Start

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
	// Initialize client
	client, err := bbb.NewClient(
		"https://domain/bigbluebutton",
		"secret-key",
		bbb.WithTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}	

	fmt.Printf("BigBlueButton Create Meeting API...\n")
	meeting , err := client.CreateMeeting(context.Background(), &requests.CreateMeetingRequest{
		Name:            "Team Meeting",
		MeetingID:       "meeting-123",
		AttendeePW:      "ap",
		ModeratorPW:     "mp",
		Welcome:         "Welcome to our team meeting!",
		Record:          true,
		MaxParticipants: 50,
		Meta: map[string]string{
			"meeting-purpose": "weekly-sync",
		},
	})
	if err != nil {
		log.Fatalf("Failed to create meeting: %v", err)
	}
    fmt.Printf("Meeting created... \n"
    fmt.Printf("Meeting ID: %s\n", meeting.MeetingID)
	fmt.Printf("BigBlueButton Join Meeting API...\n")
	joinURL, err := client.JoinMeeting(context.Background(), &requests.JoinMeetingRequest{
		MeetingID: meeting.MeetingID,
		Password:  "mp", // Moderator password (use "ap" for attendees)
		FullName:  "John Doe",
		UserID:    "user-123",
		UserData: map[string]string{
			"role": "moderator",
		},
	})
	fmt.Printf("Join Url: %s\n", joinURL)
	if err != nil {
		log.Fatalf("Failed to generate join URL: %v", err)
	}
}

```

## API Coverage

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
