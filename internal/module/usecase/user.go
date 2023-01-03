package usecase

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/deall-users/internal/model"
	"github.com/deall-users/internal/module/repository"
	e "github.com/deall-users/pkg/error"
	"github.com/deall-users/pkg/utils"
)

func NewUserUsecase(urp repository.UserRepository) UserUsecase {
	return &userUsecase{urp}
}

type userUsecase struct {
	userRp repository.UserRepository
}

func (uc userUsecase) FindAllUsers(ctx context.Context, pagination model.Pagination) (int, []*model.User, error) {
	criteria := make(map[string]interface{})

	users, err := uc.userRp.FindAllUsers(ctx, pagination, criteria)
	if err != nil {
		return 0, nil, err
	}

	total, err := uc.userRp.CountUsers(ctx, criteria)
	if err != nil {
		return 0, nil, err
	}

	return total, users, nil
}

func (uc userUsecase) FindUserByID(ctx context.Context, userID uint64) (*model.User, error) {
	return uc.userRp.FindUserByID(ctx, userID)
}

func (uc userUsecase) CreateUser(ctx context.Context, req model.CreateUser) (*model.User, error) {
	pass, err := utils.Decrypt(req.Password)
	if err != nil {
		return nil, e.WrapErrorf(err, http.StatusBadRequest, "Invalid password format")
	}

	user := model.User{
		RoleID:   req.RoleID,
		Username: strings.ToLower(req.Username),
		Fullname: strings.Trim(req.Fullname, " "),
		Email:    strings.ToLower(req.Email),
		Password: utils.HashText(pass),
	}

	err = uc.userRp.CreateUser(ctx, &user)
	if err != nil {
		if errors.Is(err, e.ErrDataAlreadyExist) {
			return nil, e.NewErrorf(http.StatusBadRequest, "username has been taken")
		}
		return nil, err
	}

	newUser, err := uc.userRp.FindUserByID(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (uc userUsecase) UpdateUser(ctx context.Context, req model.UpdateUser) (*model.User, error) {
	user := model.User{
		UserID:   req.UserID,
		RoleID:   req.RoleID,
		Username: strings.ToLower(req.Username),
		Fullname: strings.Trim(req.Fullname, " "),
		Email:    strings.ToLower(req.Email),
	}

	err := uc.userRp.UpdateUser(ctx, &user)
	if err != nil {
		if errors.Is(err, e.ErrDataAlreadyExist) {
			return nil, e.NewErrorf(http.StatusBadRequest, "username has been taken")
		}
		return nil, err
	}

	updatedUser, err := uc.userRp.FindUserByID(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil

}

func (uc userUsecase) DeleteUser(ctx context.Context, userID uint64) error {
	return uc.userRp.DeleteUser(ctx, &model.User{UserID: userID})
}

func (uc userUsecase) FindRoleByID(ctx context.Context, roleID uint64) (*model.Role, error) {
	return uc.userRp.FindRoleByID(ctx, roleID)
}
