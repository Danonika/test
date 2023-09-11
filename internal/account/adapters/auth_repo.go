package adapters

import (
	"context"

	"github.com/jackc/pgx"

	"github.com/danonika/test/pkg/postgres"

	"github.com/danonika/test/internal/account/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthorizationRepositoryPostgre struct {
	pg *pgxpool.Pool
}

var _ domain.AuthorizationRepository = (*AuthorizationRepositoryPostgre)(nil)

func NewAuthorizationRepositoryPostgre(pg *pgxpool.Pool) *AuthorizationRepositoryPostgre {
	return &AuthorizationRepositoryPostgre{
		pg: pg,
	}
}

func (r *AuthorizationRepositoryPostgre) Create(ctx context.Context, input *domain.SignUpInput) error {
	err := postgres.InTransaction(ctx, r.pg, func(ctx context.Context, tx pgx.Tx) error {
		if err := r.create(ctx, tx, input); err != nil {
			return nil
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorizationRepositoryPostgre) create(ctx context.Context, tx pgx.Tx, input *domain.CreateInput) error {
	const sql = `
		insert into "accounts" 
			(Name, balance)
		values
			($1, $2)
	`
	if _, err := tx.Exec(ctx, sql, input.Name, input.Balance); err != nil {
		return err
	}

	return nil
}
