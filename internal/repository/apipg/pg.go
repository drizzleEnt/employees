package apipg

import (
	"context"
	"fmt"

	"github.com/drizzleent/emplyees/internal/client/db"
	"github.com/drizzleent/emplyees/internal/model"
	"github.com/drizzleent/emplyees/internal/repository"
	"github.com/georgysavva/scany/pgxscan"
)

const (
	idColumn = "employee.id"

	employeeTable         = "employee"
	employeeNameColumn    = "name"
	employeeSurnameColumn = "surname"
	employeePhoneColumn   = "phone"
	employeeCompanyColumn = "companyid"
	employeeIdColumn      = "employee_id"

	passportTable    = "passport"
	typeColumn       = "type"
	numberColumn     = "number"
	passportIdColumn = "passport.employee_id"

	departamentTable = "departament"
	depNameColumn    = "depname"
	depPhoneColumn   = "depphone"
	depIdColumn      = "departament.employee_id"
	depCompanyId     = "company_id"
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

	quary = fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) values ($1, $2, $3, $4)",
		departamentTable, employeeIdColumn, depNameColumn, depPhoneColumn, depCompanyId)

	q = db.Quary{
		Name:     "repository.pg.create.deprtamentTable",
		QuaryRow: quary,
	}
	args = []interface{}{id, employee.Departament.Name, employee.Departament.Phone, employee.CompanyId}
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
		passportTable, typeColumn, numberColumn, passportIdColumn)

	q = db.Quary{
		Name:     "repository.pg.update.passportTable",
		QuaryRow: quary,
	}

	args = []interface{}{employee.Passport.Type, employee.Passport.Number, employee.Id}

	res, err = r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to update employee passport: %v, tag: %v", err, res)
	}

	quary = fmt.Sprintf("UPDATE %s SET %s=$1, %s=$2, %s=$3 WHERE %s=$4",
		departamentTable, depNameColumn, depPhoneColumn, depCompanyId, depIdColumn)

	q = db.Quary{
		Name:     "repository.pg.update.departamentTable",
		QuaryRow: quary,
	}

	args = []interface{}{employee.Departament.Name, employee.Departament.Phone, employee.CompanyId, employee.Id}

	res, err = r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to update departament: %v, tag: %v", err, res)
	}

	return nil
}
func (r *repo) GetWithCompany(ctx context.Context, companyId int) ([]model.Employee, error) {
	quary := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s JOIN %s ON %s = %s JOIN %s ON %s = %s WHERE %s = $1",
		employeeNameColumn, employeeSurnameColumn, employeePhoneColumn, employeeCompanyColumn, typeColumn, numberColumn, depNameColumn, depPhoneColumn,
		employeeTable, passportTable, passportIdColumn, idColumn, departamentTable, idColumn, depIdColumn, employeeCompanyColumn)

	q := db.Quary{
		Name:     "repository.pg.get_with_company.employeeTable",
		QuaryRow: quary,
	}

	agrs := []interface{}{companyId}
	rows, err := r.db.DB().QuaryContext(ctx, q, agrs...)
	if err != nil {
		return nil, err
	}
	res := make([]model.Employee, 1)

	err = pgxscan.ScanAll(&res, rows)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (r *repo) GetWithDepartament(ctx context.Context, dep string, id int) ([]model.Employee, error) {

	quary := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s JOIN %s ON %s = %s JOIN %s ON %s = %s WHERE %s = $1 AND %s = $2",
		employeeNameColumn, employeeSurnameColumn, employeePhoneColumn, employeeCompanyColumn, typeColumn, numberColumn, depNameColumn, depPhoneColumn,
		employeeTable, passportTable, passportIdColumn, idColumn, departamentTable, idColumn, depIdColumn, depNameColumn, employeeCompanyColumn)
	q := db.Quary{
		Name:     "repository.pg.get_with_departament.employeeTable",
		QuaryRow: quary,
	}

	args := []interface{}{dep, id}
	rows, err := r.db.DB().QuaryContext(ctx, q, args...)

	if err != nil {
		return nil, err
	}

	res := make([]model.Employee, 1)

	err = pgxscan.ScanAll(&res, rows)
	if err != nil {
		return nil, err
	}

	return res, nil
}
