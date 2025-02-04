package loadtest

import (
	"stress-test/internal/domain"
	"stress-test/pkg/httploader"
)

type Service struct {
	repo     domain.ResultRepository
	loader   *httploader.Loader
	config   domain.TestConfig
}

func NewService(repo domain.ResultRepository, config domain.TestConfig) *Service {
	return &Service{
		repo:   repo,
		loader: httploader.NewLoader(),
		config: config,
	}
}

func (s *Service) RunTest() error {
	result := s.loader.SendRequests(s.config.URL, s.config.RequestsPerSec)
	s.repo.Store(result)
	return nil
}

func (s *Service) GetResults() []domain.TestResult {
	return s.repo.GetAll()
} 