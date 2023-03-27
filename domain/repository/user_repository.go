package repository

import (
	"context"

	"github.com/yoshikawa-river/ChatApp/domain/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, u *entity.User) (*entity.User, error)
	Validate(ctx context.Context, u *entity.User) error
}
