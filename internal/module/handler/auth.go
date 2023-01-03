package handler

import (
	"net/http"

	"github.com/deall-users/internal/model"
	uc "github.com/deall-users/internal/module/usecase"
	e "github.com/deall-users/pkg/error"
	res "github.com/deall-users/pkg/response"
	"github.com/deall-users/pkg/utils"
	"github.com/gin-gonic/gin"
)

func NewAuthHandler(aUc uc.AuthUsecase) AuthHandler {
	return &authHandler{aUc}
}

type authHandler struct {
	authUc uc.AuthUsecase
}

func (h authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	//bind body json
	reqLogin := model.Login{}
	errV, err := e.BindValidateJSON(c, &reqLogin)
	if len(errV) > 0 {
		e.RespErrValidation(errV).RespJSON(c)
		return
	}

	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	token, err := h.authUc.Login(ctx, reqLogin)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(token).RespJSON(c)
}

func (h authHandler) Logout(c *gin.Context) {
	ctx := c.Request.Context()

	err := h.authUc.Logout(ctx)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(nil).RespJSON(c)
}

func (h authHandler) UserTokenValidation(c *gin.Context) {
	ctx := c.Request.Context()

	err := h.authUc.UserTokenValidation(ctx)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(nil).RespJSON(c)
}

func (h authHandler) UserRefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	token, err := h.authUc.UserRefreshToken(ctx)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(token).RespJSON(c)
}

func (h authHandler) Encryption(c *gin.Context) {
	payload := c.Query("payload")
	if payload == "" {
		res.DefaultResponse("payload required on param query", nil, nil, http.StatusBadRequest).RespJSON(c)
		return
	}
	enc, _ := utils.Encrypt(payload)

	res.RespSuccess(enc).RespJSON(c)
}

func (h authHandler) Decryption(c *gin.Context) {
	payload := c.Query("payload")
	if payload == "" {
		res.DefaultResponse("payload required on param query", nil, nil, http.StatusBadRequest).RespJSON(c)
		return
	}
	dec, _ := utils.Decrypt(payload)

	res.RespSuccess(dec).RespJSON(c)
}
