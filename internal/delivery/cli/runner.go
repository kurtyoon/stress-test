package cli

import (
	"fmt"
	"time"
	"stress-test/internal/domain"
	"stress-test/internal/usecase/loadtest"
)

type Runner struct {
	service *loadtest.Service
	config  domain.TestConfig
}

func NewRunner(service *loadtest.Service, config domain.TestConfig) *Runner {
	return &Runner{
		service: service,
		config:  config,
	}
}

func (r *Runner) Run() {
	for i := 0; i < r.config.Duration; i++ {
		fmt.Printf("\n[INFO] Running %d seconds, %d requests per second", i+1, r.config.RequestsPerSec)
		r.service.RunTest()
		time.Sleep(1 * time.Second)
	}
	fmt.Println("\n[INFO] Test completed")
} 