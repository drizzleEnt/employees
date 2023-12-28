package employe

import (
	"context"

	"github.com/drizzleent/emplyees/internal/model"
)

func (s *Service) GetWithDepartament(ctx context.Context, dep string, id int) ([]model.Employee, error) {

	res, err := s.repository.GetWithDepartament(ctx, dep, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
