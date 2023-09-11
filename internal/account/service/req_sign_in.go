package service

import (
	"context"

	"github.com/danonika/test/internal/account/domain"
)

type CreateRequest struct {
	Name    string
	Balance int
}

func (s *service) Create(ctx context.Context, input *CreateRequest) error {
	err := s.authRepo.Create(ctx, &domain.CreateInput{})
	if err != nil {
		return err
	}
	return err
}
