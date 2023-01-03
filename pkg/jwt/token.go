package jwt

import (
	cfg "github.com/deall-users/config"
	errors "github.com/deall-users/pkg/error"
	"github.com/deall-users/pkg/utils"

	"github.com/golang-jwt/jwt"
)

//JWT is adopted to this token. Authorizaton type is Bearer.
type UserToken struct {
	UserID uint64 `json:"user_id"`
	RoleID uint64 `json:"role_id"`
	jwt.StandardClaims
}

func getToken(data *UserToken, isRefresh bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	key := cfg.JWT_ACCESS_KEY
	if isRefresh {
		key = cfg.JWT_REFRESH_KEY
	}
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return errors.ErrUnknownError.Error(), err
	}

	enc, err := utils.Encrypt(tokenString)
	if err != nil {
		return "", err
	}

	return enc, nil
}

func validateToken(tokenString string, isRefresh bool) (token *jwt.Token, err error) {
	if len(tokenString) <= 0 {
		return nil, errors.ErrInvalidToken
	}

	key := cfg.JWT_ACCESS_KEY
	if isRefresh {
		key = cfg.JWT_REFRESH_KEY
	}

	token, err = jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err == nil {
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.ErrInvalidToken //That's not even a token
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.ErrTokenExipred
		} else {
			return nil, errors.ErrInvalidToken //Couldn't handle this token
		}
	} else {
		return nil, errors.ErrUnknownError
	}
}
