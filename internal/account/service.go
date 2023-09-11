package account

import (
	"context"

	"github.com/danonika/test/internal/account/adapters"

	"github.com/danonika/test/internal/account/service"
	"github.com/danonika/test/pkg/jwt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func MakeService(ctx context.Context, opts ...Option) service.Service {
	deps := &dependencies{}
	deps.setDefaults()
	for _, opt := range opts {
		opt(deps)
	}
	authRepo := adapters.NewAuthorizationRepositoryPostgre(deps.pg)
	svc := service.NewService(authRepo)
	return svc

}

type Option func(*dependencies)

type dependencies struct {
	pg     *pgxpool.Pool
	jwtcli *jwt.Client
}

func WithPostgre(pg *pgxpool.Pool) Option {
	return func(d *dependencies) {
		d.pg = pg
	}
}

func WithJWT(jwtcli *jwt.Client) Option {
	return func(d *dependencies) {
		d.jwtcli = jwtcli
	}
}

func (d *dependencies) setDefaults() {
	// pass
}
