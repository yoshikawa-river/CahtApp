package infra

import (
	"context"
	"database/sql"
	"errors"
	"net/mail"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/yoshikawa-river/ChatApp/domain/entity"
	"github.com/yoshikawa-river/ChatApp/domain/repository"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(ctx context.Context, u *entity.User) (*entity.User, error) {
	if err := ur.Validate(ctx, u); err != nil {
		return nil, err
	}
	return u, u.Insert(ctx, ur.DB, boil.Infer())
}

func (ur *UserRepository) Validate(ctx context.Context, u *entity.User) error {
	if len(u.Name) == 0 {
		return errors.New("The name field is required.")
	}

	if len(u.Name) >= 30 {
		return errors.New("The name must not be greater than 30")
	}

	if len(u.Email) == 0 {
		return errors.New("The Email field is required.")
	}

	if len(u.Email) >= 319 {
		return errors.New("The email must not be greater than 319")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("The email must be a valid email address.")
	}

	if exists, _ := entity.Users(qm.Where("email=?", u.Email)).Exists(ctx, ur.DB); exists {
		return errors.New("The email has already been taken.")
	}

	return nil
}
