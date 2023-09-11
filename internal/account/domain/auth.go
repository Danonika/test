package domain

import "context"

type CreateInput struct {
	Name    string
	Balance int
}

type AccountRepository interface {
	Create(ctx context.Context, input *CreateInput) error
}
