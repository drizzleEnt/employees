package service

import (
	"context"

	"github.com/drizzleent/emplyees/internal/model"
)

type ApiService interface {
	Create(ctx context.Context, employee *model.Employee) (int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, employee *model.Employee) error
	GetWithCompany(ctx context.Context, companyId int) ([]*model.Employee, error)
	GetWithDepartament(ctx context.Context, companyId int, departmentId int) ([]*model.Employee, error)
}
