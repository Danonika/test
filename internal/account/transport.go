package account

import (
	"net/http"

	"github.com/danonika/test/internal/common"
	"github.com/go-kit/kit/transport"
	"github.com/gorilla/mux"

	accountEndpoint "github.com/danonika/test/internal/account/endpoint"
	accountService "github.com/danonika/test/internal/account/service"
	accountTranport "github.com/danonika/test/internal/account/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
)

func MakeHandler(svc accountService.Service, logger kitlog.Logger) http.Handler {
	var (
		opts = []kithttp.ServiceOption{
			kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
			kithttp.ServerErrorEncoder(common.EncodeError),
			kithttp.ServerBefore(accountTranport.WithRequestToken),
		}
		// mdlwChain = endpoint.Chain(
		// 	authEndpoint.MakeClaimsMiddleware[authDomain.UserClaims](jwtcli),
		// )
		createHandler = kithttp.NewServer(
			accountEndpoint.MakeCreateEndpoint(svc),
			accountTranport.DecodeCreateRequest, common.EncodeResponse,
			opts...,
		)
	)

	r := mux.NewRouter()
	r.Handle("/account/create", createHandler).Methods("GET")
	return r
}
