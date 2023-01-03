package jwt

import (
	"time"

	cfg "github.com/deall-users/config"
	"github.com/deall-users/internal/model"
	errors "github.com/deall-users/pkg/error"
	"github.com/deall-users/pkg/utils"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user *model.User, subject string, isRefresh bool) (token string, expiredAt *time.Time, err error) {
	today := utils.TimeNow()
	duration := cfg.JWT_ACCESS_EXPIRATION
	if isRefresh {
		duration = cfg.JWT_REFRESH_EXPIRATION
	}

	expiration := today.Add(time.Hour * time.Duration(duration))
	tokenData := &UserToken{
		UserID: user.UserID,
		RoleID: user.RoleID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    cfg.APP_NAME,
			IssuedAt:  today.Unix(),
			ExpiresAt: expiration.Unix(),
			Subject:   subject,
		},
	}

	token, err = getToken(tokenData, isRefresh)
	if err != nil {
		return "", nil, err
	}

	return token, &expiration, nil
}

func ValidationAndExtraction(tokenString string, isRefresh bool) (*UserToken, error) {
	dec, err := utils.Decrypt(tokenString)
	if err != nil {
		return nil, errors.ErrInvalidToken
	}

	token, err := validateToken(dec, isRefresh)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserToken)
	if ok && token.Valid && claims != nil {
		return claims, nil
	}

	return nil, errors.ErrInvalidToken
}
