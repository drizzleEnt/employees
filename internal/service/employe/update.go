package employe

import (
	"context"

	"github.com/drizzleent/emplyees/internal/model"
)

func (s *Service) Update(ctx context.Context, employee *model.Employee) error {
	err := s.repository.Update(ctx, employee)

	if err != nil {
		return err
	}

	return nil
}
