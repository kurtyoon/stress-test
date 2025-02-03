# HTTP Load Testing Tool

A simple Go-based HTTP load testing tool for performing stress tests on web endpoints.

## Features

- Performs HTTP GET requests to specified URLs
- Configurable requests per second (RPS)
- Adjustable test duration
- Provides statistics by HTTP status codes
- Measures average response latency
- Concurrent request handling using goroutines

## Installation

Clone the repository and navigate to the project directory:

```sh
git clone [repository-url]
cd [repository-name]
```

## Usage

Run the program with the following command:

```sh
go run main.go -url [TARGET_URL] -rps [REQUESTS_PER_SECOND] -duration [DURATION_IN_SECONDS]
```

### Required Parameters

- `-url`: Target URL to test
- `-rps`: Number of requests per second
- `-duration`: Test duration in seconds

### Example

```sh
go run main.go -url https://example.com -rps 10 -duration 5
```

## Output

The program provides the following metrics:

- Total number of requests
- Successful responses (HTTP 200)
- Error responses (5xx)
- Average response latency
- Test duration
- Summary of HTTP status codes

Example output:

```sh
=== Results ===
Total Requests: 10
Success Response (200): 10
Error Response (5xx): 0
Average Latency: 41.604299ms
Duration: 53.27525ms

=== HTTP Status Code Summary ===
HTTP Status 200: 10
```

## Caution

- Consider the target server's capacity and network conditions when setting RPS
- Use with caution in production environments
- Be mindful of the load you're generating on the target system
