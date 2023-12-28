package employe

import (
	"context"

	"github.com/drizzleent/emplyees/internal/model"
)

func (s *Service) GetWithCompany(ctx context.Context, companyId int) ([]model.Employee, error) {

	res, err := s.repository.GetWithCompany(ctx, companyId)

	if err != nil {
		return nil, err
	}

	return res, nil
}
