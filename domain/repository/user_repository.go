package repository

import (
	"context"

	"github.com/yoshikawa-river/ChatApp/domain/models"
)

type IUserRepository interface {
	Create(ctx context.Context, u *models.User) (*models.User, error)
	Validate(ctx context.Context, u *models.User) error
}
