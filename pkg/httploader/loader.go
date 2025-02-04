package httploader

import (
	"net/http"
	"sync"
	"time"
	"stress-test/internal/domain"
)

type Loader struct {
	client *http.Client
}

func NewLoader() *Loader {
	return &Loader{
		client: &http.Client{},
	}
}

func (l *Loader) SendRequests(url string, count int) domain.TestResult {
	var wg sync.WaitGroup
	results := make(chan int, count)
	latency := make(chan time.Duration, count)
	
	startTime := time.Now()

	for i := 0; i < count; i++ {
		wg.Add(1)
		go l.sendRequest(&wg, url, results, latency)
	}

	wg.Wait()
	close(results)
	close(latency)

	return l.processResults(results, latency, startTime)
}

func (l *Loader) sendRequest(wg *sync.WaitGroup, url string, results chan<- int, latency chan<- time.Duration) {
	defer wg.Done()

	start := time.Now()
	resp, err := l.client.Get(url)
	duration := time.Since(start)

	if err != nil {
		results <- 0
		return
	}
	
	results <- resp.StatusCode
	latency <- duration
	resp.Body.Close()
}

func (l *Loader) processResults(results <-chan int, latency <-chan time.Duration, startTime time.Time) domain.TestResult {
	successCount, errorCount := 0, 0
	statusCounts := make(map[int]int)
	var totalLatency time.Duration

	for res := range results {
		if res == 200 {
			successCount++
		} else {
			errorCount++
		}
		statusCounts[res]++
	}

	for lat := range latency {
		totalLatency += lat
	}

	totalRequests := successCount + errorCount

	return domain.TestResult{
		Timestamp:      time.Now(),
		TotalRequests:  totalRequests,
		SuccessCount:   successCount,
		ErrorCount:     errorCount,
		AverageLatency: totalLatency / time.Duration(totalRequests),
		Duration:       time.Since(startTime),
		StatusCounts:   statusCounts,
	}
} 