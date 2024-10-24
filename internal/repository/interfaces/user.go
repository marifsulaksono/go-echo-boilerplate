package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context) (*[]model.User, error)
	GetById(ctx context.Context, id uuid.UUID) (data *model.User, err error)
	GetByEmail(ctx context.Context, email string) (data *model.User, err error)
	Create(ctx context.Context, payload *model.User) (data *model.UserResponse, err error)
	Update(ctx context.Context, payload *model.User, id uuid.UUID) (data *model.UserResponse, err error)
	Delete(ctx context.Context, id uuid.UUID) error
}
