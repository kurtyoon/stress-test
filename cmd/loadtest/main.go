package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"stress-test/internal/delivery/cli"
	delivery "stress-test/internal/delivery/http"
	"stress-test/internal/domain"
	"stress-test/internal/repository/memory"
	"stress-test/internal/usecase/loadtest"
)

func main() {
	config := parseFlags()
	
	repo := memory.NewResultRepository()
	service := loadtest.NewService(repo, config)
	handler := delivery.NewHandler(service)
	runner := cli.NewRunner(service, config)

	handler.RegisterRoutes()
	go http.ListenAndServe(":8080", nil)
	fmt.Println("[INFO] Dashboard available at http://localhost:8080")

	runner.Run()
	
	select {}
}

func parseFlags() domain.TestConfig {
	var config domain.TestConfig
	
	flag.StringVar(&config.URL, "url", "", "Target URL")
	flag.IntVar(&config.RequestsPerSec, "rps", 0, "Requests per second")
	flag.IntVar(&config.Duration, "duration", 0, "Test duration in seconds")
	flag.Parse()

	if len(os.Args) == 1 {
		printUsage()
		os.Exit(1)
	}

	if config.URL == "" || config.RequestsPerSec == 0 || config.Duration == 0 {
		fmt.Println("Error: url, rps, and duration are required")
		os.Exit(1)
	}

	return config
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  -url string")
	fmt.Println("        Target URL")
	fmt.Println("  -rps int")
	fmt.Println("        Requests per second")
	fmt.Println("  -duration int")
	fmt.Println("        Test duration in seconds")
} 