package endpoint

import (
	"context"

	"github.com/danonika/test/internal/account/service"
	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	Name    string
	Balance int
}

type CreateResponce struct {
	Err error `json:"err,omitempty"`
}

func (r *CreateResponce) Error() error { return r.Err }

func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) error {
		req := request.(CreateRequest)
		err := s.Create(ctx, &service.CreateRequest{
			Name:    req.Name,
			Balance: req.Balance,
		})
		return err
	}
}
