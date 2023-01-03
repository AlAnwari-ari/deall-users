package usecase

import (
	"context"
	"net/http"

	"github.com/deall-users/internal/model"
	"github.com/deall-users/internal/module/repository"
	e "github.com/deall-users/pkg/error"
	"github.com/deall-users/pkg/jwt"
	"github.com/deall-users/pkg/utils"
)

func NewAuthUsecase(urp repository.UserRepository) AuthUsecase {
	return &authUsecase{urp}
}

type authUsecase struct {
	userRp repository.UserRepository
}

func (uc authUsecase) Login(ctx context.Context, req model.Login) (*model.TokenResponse, error) {
	//check user exist or not
	user, err := uc.userRp.FindSingleUser(ctx, map[string]interface{}{"username": req.Username})
	if err != nil {
		return nil, err
	}

	//compare password
	pass, err := utils.Decrypt(req.Password)
	if err != nil || !utils.CompareHashPassword(user.Password, pass) {
		return nil, e.ErrAuthLogin
	}

	//generate access token
	token, expired, err := jwt.GenerateToken(user, "Login", false)
	if err != nil {
		return nil, e.WrapErrorf(err, http.StatusInternalServerError, "token")
	}

	//generate refresh token
	tokenRef, expiredRef, err := jwt.GenerateToken(user, "Refresh Login", true)
	if err != nil {
		return nil, e.WrapErrorf(err, http.StatusInternalServerError, "refresh token")
	}

	return &model.TokenResponse{
		AccessToken:    token,
		RefreshToken:   tokenRef,
		AccessExpired:  expired,
		RefreshExpired: expiredRef,
	}, nil
}

func (uc authUsecase) Logout(ctx context.Context) (err error) {
	_, ok := ctx.Value(model.TokenCtxKey).(string)
	if !ok {
		return e.ErrInvalidToken
	}

	return nil
}

func (uc authUsecase) UserTokenValidation(ctx context.Context) (err error) {
	_, ok := ctx.Value(model.UserIDCtxKey).(uint64)
	if !ok {
		return e.ErrInvalidToken
	}

	return nil
}

func (uc authUsecase) UserRefreshToken(ctx context.Context) (*model.TokenResponse, error) {
	userID, ok := ctx.Value(model.UserIDCtxKey).(uint64)
	if !ok {
		return nil, e.ErrInvalidToken
	}

	// get newest data user
	user, err := uc.userRp.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	//generate access token
	token, expired, err := jwt.GenerateToken(user, "ReLogin", false)
	if err != nil {
		return nil, e.WrapErrorf(err, http.StatusInternalServerError, "token")
	}

	//generate refresh token
	tokenRef, expiredRef, err := jwt.GenerateToken(user, "Refresh ReLogin", true)
	if err != nil {
		return nil, e.WrapErrorf(err, http.StatusInternalServerError, "refresh token")
	}

	return &model.TokenResponse{
		AccessToken:    token,
		RefreshToken:   tokenRef,
		AccessExpired:  expired,
		RefreshExpired: expiredRef,
	}, nil
}
