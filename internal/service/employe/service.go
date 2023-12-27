package employe

import (
	"github.com/drizzleent/emplyees/internal/repository"
	"github.com/drizzleent/emplyees/internal/service"
)

type Service struct {
	repository repository.ApiRepository
}

func NewService(repo repository.ApiRepository) service.ApiService {
	return &Service{
		repository: repo,
	}
}
