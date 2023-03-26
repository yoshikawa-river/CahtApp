package usecase

import (
	"context"

	"github.com/yoshikawa-river/ChatApp/domain/repository"
	"github.com/yoshikawa-river/ChatApp/usecase/models"
)

type IUserInteractor interface {
	Create(ctx context.Context, u *models.User) (*models.User, error)
}

type UserInteractor struct {
	ur repository.IUserRepository
}

func NewUserInteractor(userRepository repository.IUserRepository) IUserInteractor {
	return &UserInteractor{
		ur: userRepository,
	}
}

func (uu *UserInteractor) Create(ctx context.Context, u *models.User) (*models.User, error) {
	user := models.UserToDomainModel(u)
	du, err := uu.ur.Create(ctx, user)
	if err != nil {
		panic(err)
	}
	return models.UserFromDomainModel(du), err
}
