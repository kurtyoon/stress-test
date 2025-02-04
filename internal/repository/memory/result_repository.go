package memory

import (
	"sync"

	"stress-test/internal/domain"
)

type ResultRepository struct {
	results []domain.TestResult
	mu      sync.RWMutex
}

func NewResultRepository() *ResultRepository {
	return &ResultRepository{
		results: make([]domain.TestResult, 0),
	}
}

func (r *ResultRepository) Store(result domain.TestResult) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.results = append(r.results, result)
}

func (r *ResultRepository) GetAll() []domain.TestResult {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.results
} 