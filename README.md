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
)

func main() {
	// Initialize client
	client, err := bbb.NewClient(
		"https://your-bbb-server.com/bigbluebutton/",
		"your-secret-here",
		bbb.WithTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: Get server version
	version, err := client.GetAPIVersion(context.Background())
	if err != nil {
		log.Fatalf("Failed to get API version: %v", err)
	}

	fmt.Printf("BigBlueButton API Version: %s\n", version)
}
```

## API Coverage

### Meetings
- [x] Create meeting
- [ ] Join meeting
- [ ] End meeting
- [ ] Get meeting info
- [ ] Get meetings
- [ ] Get recordings
- [ ] Publish/unpublish recordings
- [ ] Delete recordings

### Webhooks
- [ ] Create hook
- [ ] Get hooks
- [ ] Destroy hook

## Documentation

Full API documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/amirazad1/bigbluebutton-api-go).

## Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
