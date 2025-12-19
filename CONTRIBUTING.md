# Contributing to BigBlueButton Go Client

Thank you for your interest in contributing to the BigBlueButton Go Client! We appreciate your time and effort in making this project better.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Code Style](#code-style)
- [Testing](#testing)
- [Pull Requests](#pull-requests)
- [Reporting Issues](#reporting-issues)
- [Feature Requests](#feature-requests)
- [License](#license)

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
   ```bash
   git clone https://github.com/your-username/bigbluebutton-api-go.git
   cd bigbluebutton-api-go
   ```
3. Install dependencies (no external dependencies required)
4. Create a feature branch
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Workflow

1. Make your changes
2. Run tests and ensure they pass
   ```bash
   go test ./... -v -cover
   ```
3. Format your code
   ```bash
   go fmt ./...
   ```
4. Lint your code
   ```bash
   golangci-lint run
   ```
5. Commit your changes with a descriptive commit message
   ```bash
   git commit -m "feat: add new feature"
   ```
6. Push to your fork
   ```bash
   git push origin feature/your-feature-name
   ```
7. Open a pull request

## Code Style

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Keep functions small and focused on a single responsibility
- Write clear and concise comments for exported types and functions
- Use meaningful variable and function names

## Testing

- Write tests for new features and bug fixes
- Ensure all tests pass before submitting a pull request
- Maintain or improve test coverage
- Use table-driven tests when appropriate

## Pull Requests

1. Keep pull requests focused on a single feature or bug fix
2. Include tests for new features and bug fixes
3. Update documentation as needed
4. Ensure the CI build passes
5. Request a review from one of the maintainers

## Reporting Issues

When reporting issues, please include:

- A clear and descriptive title
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Version of the library
- Version of Go
- Any relevant error messages or logs

## Feature Requests

We welcome feature requests! Please open an issue and describe:

- The feature you'd like to see
- Why this feature would be useful
- Any potential implementation ideas

## License

By contributing to this project, you agree that your contributions will be licensed under the [MIT License](LICENSE).
