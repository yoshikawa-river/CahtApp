package usecase

import (
	"context"

	"github.com/yoshikawa-river/ChatApp/domain/models"
	"github.com/yoshikawa-river/ChatApp/domain/repository"
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

func (ui *UserInteractor) Create(ctx context.Context, u *models.User) (*models.User, error) {
	user := models.UserToDomainModel(u)
	eu, err := ui.ur.Create(ctx, user)
	if err != nil {
		panic(err)
	}
	return models.UserFromDomainModel(eu), err
}

// ユーザー更新の関数を追加
func (ui *UserInteractor) Update(ctx context.Context, u *models.User) (*models.User, error) {
	user := models.UserToDomainModel(u)
	eu, err := ui.ur.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return models.UserFromDomainModel(eu), nil
}
