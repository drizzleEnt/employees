package apipg

import (
	"context"
	"fmt"

	"github.com/drizzleent/emplyees/internal/client/db"
	"github.com/drizzleent/emplyees/internal/model"
	"github.com/drizzleent/emplyees/internal/repository"
)

const (
	idColumn = "id"

	employeeTable         = "employee"
	employeeNameColumn    = "name"
	employeeSurnameColumn = "surname"
	employeePhoneColumn   = "phone"
	employeeCompanyColumn = "companyid"
	employeeIdColumn      = "employee_id"

	passportTable = "passport"
	typeColumn    = "type"
	numberColumn  = "number"

	departamentTable = "departament"
	depNameColumn    = "name"
	depPhoneColumn   = "phone"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ApiRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, employee *model.Employee) (int, error) {
	quary := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) values ($1, $2, $3, $4) RETURNING id",
		employeeTable, employeeNameColumn, employeeSurnameColumn, employeePhoneColumn, employeeCompanyColumn)

	q := db.Quary{
		Name:     "repository.pg.create.employeeTable",
		QuaryRow: quary,
	}
	args := []interface{}{employee.Name, employee.Surname, employee.Phone, employee.CompanyId}
	var id int

	err := r.db.DB().QuaryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	quary = fmt.Sprintf("INSERT INTO %s (%s, %s, %s) values ($1, $2, $3)",
		passportTable, employeeIdColumn, typeColumn, numberColumn)

	q = db.Quary{
		Name:     "repository.pg.create.passportTable",
		QuaryRow: quary,
	}
	args = []interface{}{id, employee.Passport.Type, employee.Passport.Number}
	r.db.DB().QuaryRowContext(ctx, q, args...)

	quary = fmt.Sprintf("INSERT INTO %s (%s, %s, %s) values ($1, $2, $3)",
		departamentTable, employeeIdColumn, depNameColumn, depPhoneColumn)

	q = db.Quary{
		Name:     "repository.pg.create.deprtamentTable",
		QuaryRow: quary,
	}
	args = []interface{}{id, employee.Departament.Name, employee.Departament.Phone}
	r.db.DB().QuaryRowContext(ctx, q, args...)

	return id, nil
}
func (r *repo) Delete(ctx context.Context, id int) error {
	quary := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", employeeTable, idColumn)

	q := db.Quary{
		Name:     "repository.pg.delete.employeeTable",
		QuaryRow: quary,
	}

	args := []interface{}{id}
	res, err := r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to delete employee: %v, tag: %v", err, res)
	}

	return nil
}
func (r *repo) Update(ctx context.Context, employee *model.Employee) error {
	//Предпочтительно сделать через tx manager и комиты
	quary := fmt.Sprintf("UPDATE %s SET %s=$1, %s=$2, %s=$3, %s=$4 WHERE %s=$5",
		employeeTable, employeeNameColumn, employeeSurnameColumn, employeePhoneColumn, employeeCompanyColumn, idColumn)

	q := db.Quary{
		Name:     "repository.pg.update.employeeTable",
		QuaryRow: quary,
	}

	args := []interface{}{employee.Name, employee.Surname, employee.Phone, employee.CompanyId, employee.Id}

	res, err := r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to update employee: %v, tag: %v", err, res)
	}

	quary = fmt.Sprintf("UPDATE %s SET %s=$1, %s=$2 WHERE %s=$3",
		passportTable, typeColumn, numberColumn, idColumn)

	q = db.Quary{
		Name:     "repository.pg.update.passportTable",
		QuaryRow: quary,
	}

	args = []interface{}{employee.Passport.Type, employee.Passport.Number, employee.Id}

	res, err = r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to update employee passport: %v, tag: %v", err, res)
	}

	quary = fmt.Sprintf("UPDATE %s SET %s=$1, %s=$2 WHERE %s=$3",
		departamentTable, depNameColumn, depPhoneColumn, idColumn)

	q = db.Quary{
		Name:     "repository.pg.update.departamentTable",
		QuaryRow: quary,
	}

	args = []interface{}{employee.Departament.Name, employee.Departament.Phone, employee.Id}

	res, err = r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to update departament: %v, tag: %v", err, res)
	}

	return nil
}
func (r *repo) GetWithCompany(ctx context.Context, companyId int) ([]*model.Employee, error) {
	return nil, nil
}
func (r *repo) GetWithDepartament(ctx context.Context, companyId int, departmentId int) ([]*model.Employee, error) {
	return nil, nil
}
