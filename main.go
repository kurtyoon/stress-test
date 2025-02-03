package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	url            string
	requestsPerSec int
	duration       int
)

func init() {
	flag.StringVar(&url, "url", "", "Target URL")
	flag.IntVar(&requestsPerSec, "rps", 0, "Requests per second")
	flag.IntVar(&duration, "duration", 0, "Test duration in seconds")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Usage:")
		fmt.Println("  -url string")
		fmt.Println("        Target URL")
		fmt.Println("  -rps int")
		fmt.Println("        Requests per second")
		fmt.Println("  -duration int")
		fmt.Println("        Test duration in seconds") 
		os.Exit(1)
	}

	if url == "" || requestsPerSec == 0 || duration == 0 {
		fmt.Println("Error: url, rps, and duration are required")
		os.Exit(1)
	}
}

func sendRequest(wg *sync.WaitGroup, results chan<- int, latency chan<- time.Duration) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		results <- 0
	} else {
		results <- resp.StatusCode
		latency <- duration
		resp.Body.Close()
	}
}

func performTest() {
	var wg sync.WaitGroup
	results := make(chan int, requestsPerSec)
	latency := make(chan time.Duration, requestsPerSec)

	startTime := time.Now()

	for i := 0; i < requestsPerSec; i++ {
		wg.Add(1)
		go sendRequest(&wg, results, latency)
	}

	wg.Wait()
	close(results)
	close(latency)

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
	elapsed := time.Since(startTime)

	fmt.Println("\n=== Results ===")
	fmt.Printf("Total Requests: %d\n", totalRequests)
	fmt.Printf("Success Response (200): %d\n", successCount)
	fmt.Printf("Error Response (5xx): %d\n", errorCount)
	fmt.Printf("Average Latency: %v\n", totalLatency/time.Duration(totalRequests))
	fmt.Printf("Duration: %v\n", elapsed)

	fmt.Println("\n=== HTTP Status Code Summary ===")
	for code, count := range statusCounts {
		fmt.Printf("HTTP Status %d: %d\n", code, count)
	}
}

func main() {
	for i := 0; i < duration; i++ {
		fmt.Printf("\n[INFO] Running %d seconds, %d requests per second", i+1, requestsPerSec)
		performTest()
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n[INFO] Test completed")
}
