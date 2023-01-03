package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/deall-users/internal/constanta"
	"github.com/deall-users/internal/model"
	uc "github.com/deall-users/internal/module/usecase"
	e "github.com/deall-users/pkg/error"
	"github.com/deall-users/pkg/jwt"
	res "github.com/deall-users/pkg/response"
	"github.com/deall-users/pkg/utils"
	"github.com/gin-gonic/gin"
)

func NewMiddlewareHandler(uUc uc.UserUsecase) MiddlewareHandler {
	return &middlewareHandler{uUc}
}

type middlewareHandler struct {
	userUc uc.UserUsecase
}

func (middlewareHandler) SetCors() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-CSRF-Token, Authorization, Origin, Referer, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,DELETE,PUT,POST,OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func (h middlewareHandler) TokenValidation(isRefresh bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authType, tokenHeader := utils.GetHeaderBearerToken(c)
		if constanta.Bearer != authType {
			res.DefaultResponse("Token must be Bearer type", nil, nil, http.StatusUnauthorized).RespJSON(c)
			c.Abort()
			return
		}

		if tokenHeader == "" {
			e.ErrInvalidToken.(*e.Error).RespError(nil).RespJSON(c)
			c.Abort()
			return
		}

		userAccessToken, err := jwt.ValidationAndExtraction(tokenHeader, isRefresh)
		if err != nil {
			tokenErr, ok := err.(*e.Error)
			if ok {
				tokenErr.RespError(nil).RespJSON(c)
			} else {
				e.ErrInvalidToken.(*e.Error).RespError(nil).RespJSON(c)
			}

			c.Abort()
			return
		}

		setTokenCtx := context.WithValue(c.Request.Context(), model.TokenCtxKey, tokenHeader)
		setUserIDCtx := context.WithValue(setTokenCtx, model.UserIDCtxKey, userAccessToken.UserID)
		setRoleIDCtx := context.WithValue(setUserIDCtx, model.RoleIDCtxKey, userAccessToken.RoleID)
		c.Request = c.Request.WithContext(setRoleIDCtx)
		c.Next()
	}
}

func (h middlewareHandler) Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		roleID, ok := ctx.Value(model.RoleIDCtxKey).(uint64)
		if !ok {
			e.ErrInvalidToken.(*e.Error).RespError(nil).RespJSON(c)
			c.Abort()
			return
		}

		role, err := h.userUc.FindRoleByID(ctx, roleID)
		if !ok {
			errUc, ok := err.(*e.Error)
			if ok {
				errUc.RespError(c).RespJSON(c)
			} else {
				res.DefaultResponse(err.Error(), nil, nil, http.StatusInternalServerError)
			}
			c.Abort()
			return
		}

		isAllowed := false

		switch strings.ToLower(c.Request.Method) {
		case "get":
			isAllowed = role.CanRead
		case "post":
			isAllowed = role.CanAdd
		case "put":
			isAllowed = role.CanUpdate
		case "delete":
			isAllowed = role.CanDelete
		default:
			res.DefaultResponse("Unregistered method type. Allowed: GET, POST, PUT, DELETE", nil, nil, http.StatusMethodNotAllowed)
			c.Abort()
			return
		}

		if !isAllowed {
			res.DefaultResponse("You are not allowed to access this service", nil, nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()

	}
}
