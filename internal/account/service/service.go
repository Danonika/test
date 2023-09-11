package service

import (
	"context"

	"github.com/danonika/test/internal/account/domain"
)

type Service interface {
	Create(ctx context.Context, req *CreateRequest) error
}

type service struct {
	authRepo domain.AccountRepository
}

func NewService(
	authRepo domain.AccountRepository,
) Service {
	return &service{
		authRepo: authRepo,
	}
}
