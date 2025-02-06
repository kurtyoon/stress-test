# HTTP Load Testing Tool

A simple HTTP load testing tool that performs stress tests on specified URLs with configurable requests per second and test duration.

## Project Structure

```
.
├── cmd
│   └── loadtest
│       └── main.go
├── internal
│   ├── domain
│   │   ├── result.go
│   │   └── test.go
│   ├── repository
│   │   └── memory
│   │       └── result_repository.go
│   ├── usecase
│   │   └── loadtest
│   │       └── service.go
│   └── delivery
│       ├── http
│       │   └── handler.go
│       └── cli
│           └── runner.go
├── pkg
│   └── httploader
│       └── loader.go
├── web
│   └── dashboard.html
├── Makefile
└── README.md
```

## Installation and Usage

### Build

```bash
make build
```

### Run

```bash
make run URL=https://example.com RPS=10 DURATION=5
```

### Parameters

- `URL`: Target URL to test
- `RPS`: Requests Per Second
- `DURATION`: Test duration in seconds

### System Installation

```bash
make install
```

## Docker Support

### Build Docker Image

```bash
docker build -t stress-test .
```

### Run with Docker

```bash
docker run -p 8080:8080 stress-test -url https://example.com -rps 10 -duration 5
```

## Dashboard

Test results can be monitored in real-time through the web dashboard:

- URL: http://localhost:8080
- Available Metrics:
  - Total Request Count
  - Success Rate
  - Average Response Time
  - Response Time Trend
  - HTTP Status Code Distribution
  - Recent Test Results List

## Makefile Commands

- `make build`: Build the application
- `make run`: Run the application
- `make clean`: Remove build artifacts
- `make test`: Run tests
- `make install`: Install to system
- `make help`: Display help information

## Examples

```bash
# Build and run
make build && make run URL=https://example.com RPS=10 DURATION=5

# Or run directly
make run URL=https://example.com RPS=10 DURATION=5
```

## Caution

- Consider the target server's capacity and network conditions when setting RPS
- Use with caution in production environments
- The dashboard continues to run after the test completes; use Ctrl+C to terminate the program

## Output

The program provides the following metrics:

- Total number of requests
- Successful responses (HTTP 200)
- Error responses (5xx)
- Average response latency
- Test duration
- Summary of HTTP status codes

Example output:

```
=== Results ===
Total Requests: 10
Success Response (200): 10
Error Response (5xx): 0
Average Latency: 41.604299ms
Duration: 53.27525ms

=== HTTP Status Code Summary ===
HTTP Status 200: 10
```

## Features

- Concurrent request handling using goroutines
- Real-time metrics visualization
- Configurable test parameters
- Web-based dashboard
- HTTP status code analysis
- Response time tracking
- Memory-based result storage
- Clean architecture design

## Technical Details

- Written in Go
- Uses clean architecture pattern
- Implements concurrent request handling
- Provides real-time metrics through web interface
- Supports HTTP GET requests
- Includes built-in dashboard for visualization
