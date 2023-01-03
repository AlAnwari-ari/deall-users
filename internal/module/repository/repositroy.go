package repository

import (
	"context"

	"github.com/deall-users/internal/model"
)

type UserRepository interface {
	FindAllUsers(ctx context.Context, pagination model.Pagination, criteria map[string]interface{}) ([]*model.User, error)
	FindSingleUser(ctx context.Context, criteria map[string]interface{}) (*model.User, error)
	FindUserByID(ctx context.Context, userID uint64) (*model.User, error)
	CountUsers(ctx context.Context, criteria map[string]interface{}) (int, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.User) error

	FindRoleByID(ctx context.Context, roleID uint64) (*model.Role, error)
}
