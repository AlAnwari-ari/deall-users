package usecase

import (
	"context"

	"github.com/deall-users/internal/model"
)

type UserUsecase interface {
	FindAllUsers(ctx context.Context, pagination model.Pagination) (int, []*model.User, error)
	FindUserByID(ctx context.Context, userID uint64) (*model.User, error)
	CreateUser(ctx context.Context, req model.CreateUser) (*model.User, error)
	UpdateUser(ctx context.Context, req model.UpdateUser) (*model.User, error)
	DeleteUser(ctx context.Context, userID uint64) error

	FindRoleByID(ctx context.Context, roleID uint64) (*model.Role, error)
}

type AuthUsecase interface {
	Login(ctx context.Context, req model.Login) (*model.TokenResponse, error)
	Logout(ctx context.Context) error
	UserRefreshToken(ctx context.Context) (*model.TokenResponse, error)
	UserTokenValidation(ctx context.Context) error
}
