package handler

import (
	"github.com/gin-gonic/gin"
)

type MiddlewareHandler interface {
	SetCors() gin.HandlerFunc
	TokenValidation(isRefresh bool) gin.HandlerFunc
	Permission() gin.HandlerFunc
}

type UserHandler interface {
	FindAllUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type AuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	UserTokenValidation(c *gin.Context)
	UserRefreshToken(c *gin.Context)
	Encryption(c *gin.Context)
	Decryption(c *gin.Context)
}
