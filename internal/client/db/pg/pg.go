package pg

import (
	"context"

	"github.com/drizzleent/emplyees/internal/client/db"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pg struct {
	dbc *pgxpool.Pool
}

func NewPool(dbc *pgxpool.Pool) *pg {
	return &pg{
		dbc: dbc,
	}
}

func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q db.Quary, args ...interface{}) error {
	row, err := p.QuaryContext(ctx, q, args...)

	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q db.Quary, args ...interface{}) error {
	rows, err := p.QuaryContext(ctx, q, args...)

	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) ExecContext(ctx context.Context, q db.Quary, args ...interface{}) (pgconn.CommandTag, error) {

	return p.dbc.Exec(ctx, q.QuaryRow, args...)
}

func (p *pg) QuaryContext(ctx context.Context, q db.Quary, args ...interface{}) (pgx.Rows, error) {

	return p.dbc.Query(ctx, q.QuaryRow, args...)
}

func (p *pg) QuaryRowContext(ctx context.Context, q db.Quary, args ...interface{}) pgx.Row {

	return p.dbc.QueryRow(ctx, q.QuaryRow, args...)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *pg) Close() {
	p.dbc.Close()
}
