package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/danonika/test/internal/account"
	// "github.com/danonika/test/internal/transaction"
	"github.com/danonika/test/pkg/jwt"
	"github.com/go-kit/log"
)

func main() {
	var (
		ctx              = context.Background()
		jwtsecret        = getenv("JWT_SECRET", "qaztradesecret")
		postgresLogin    = getenv("POSTGRES_LOGIN", "postgres")
		postgresPassword = getenv("POSTGRES_PASSWORD", "postgres")
		postgresHost     = getenv("POSTGRES_HOST", "localhost")
		postgresDatabase = getenv("POSTGRES_DATABASE", "finance")
		postgresURL      = fmt.Sprintf("postgresql://%s:%s@%s:5432/%s", postgresLogin, postgresPassword, postgresHost, postgresDatabase)
		jwtcli           = jwt.NewClient(jwtsecret)
	)
	pg, err := pgxpool.Connect(ctx, postgresURL)
	if err != nil {
		panic(err)
	}
	var (
		logger log.Logger
	)

	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	}
	var (
		httpLogger = log.With(logger, "component", "http")
		mux        = http.NewServeMux()
	)
	var (
		account = account.MakeService(
			ctx,
			account.WithPostgres(pg),
			account.WithJWT(jwtcli),
		)
		// transaction = transaction.MakeService(
		// 	ctx,
		// 	transaction.WithPostgres(pg),
		// )
	)
	mux.Handle("/account", account.MakeHandler(account, httpLogger))
	// mux.Handle("/transaction", transaction.MakeHandler(transaction, httpLogger))

}

func getenv(env string, fallback ...string) string {
	value := os.Getenv(env)
	if value != "" {
		return value
	}

	if len(fallback) > 0 {
		value = fallback[0]
	}
	return value
}
