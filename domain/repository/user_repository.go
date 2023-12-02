package repository

import (
	"context"

	"github.com/yoshikawa-river/ChatApp/domain/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, u *entity.User) (*entity.User, error)
	Validate(ctx context.Context, u *entity.User) error
	// 更新のインターフェイスの戻り値idを追加して実装
	Update(ctx context.Context, u *entity.User) (*entity.User, error)
}
