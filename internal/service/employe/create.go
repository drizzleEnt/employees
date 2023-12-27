package employe

import (
	"context"

	"github.com/drizzleent/emplyees/internal/model"
)

func (s *Service) Create(ctx context.Context, employee *model.Employee) (int, error) {

	id, err := s.repository.Create(ctx, employee)

	if err != nil {
		return 0, err
	}

	return id, nil
}
