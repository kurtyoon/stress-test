package domain

import "time"

type TestResult struct {
	Timestamp      time.Time      `json:"timestamp"`
	TotalRequests  int           `json:"totalRequests"`
	SuccessCount   int           `json:"successCount"`
	ErrorCount     int           `json:"errorCount"`
	AverageLatency time.Duration `json:"averageLatency"`
	Duration       time.Duration `json:"duration"`
	StatusCounts   map[int]int   `json:"statusCounts"`
}

type ResultRepository interface {
	Store(result TestResult)
	GetAll() []TestResult
} 