package handler

import (
	"github.com/deall-users/internal/model"
	uc "github.com/deall-users/internal/module/usecase"
	e "github.com/deall-users/pkg/error"
	res "github.com/deall-users/pkg/response"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(uUc uc.UserUsecase) UserHandler {
	return &userHandler{uUc}
}

type userHandler struct {
	userUc uc.UserUsecase
}

func (h userHandler) FindAllUser(c *gin.Context) {
	ctx := c.Request.Context()

	//bind pagination
	pagination := model.Pagination{}
	errV, err := e.BindValidateQuery(c, &pagination)

	if len(errV) > 0 {
		e.RespErrValidation(errV).RespJSON(c)
		return
	}

	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	total, users, err := h.userUc.FindAllUsers(ctx, pagination)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(res.NewPaginationResponse(users, pagination, total, true)).RespJSON(c)

}

func (h userHandler) FindUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	//bind uri
	userUri := model.User{}
	errV, err := e.BindValidateURI(c, &userUri)
	if len(errV) > 0 {
		e.RespErrValidation(errV).RespJSON(c)
		return
	}

	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	user, err := h.userUc.FindUserByID(ctx, userUri.UserID)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(user).RespJSON(c)
}

func (h userHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	//bind body json
	reqUser := model.CreateUser{}
	errV, err := e.BindValidateJSON(c, &reqUser)
	if len(errV) > 0 {
		e.RespErrValidation(errV).RespJSON(c)
		return
	}

	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	user, err := h.userUc.CreateUser(ctx, reqUser)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(user).RespJSON(c)
}

func (h userHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	//bind body json
	reqUser := model.UpdateUser{}
	errV, err := e.BindValidateJSON(c, &reqUser)
	if len(errV) > 0 {
		e.RespErrValidation(errV).RespJSON(c)
		return
	}

	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	user, err := h.userUc.UpdateUser(ctx, reqUser)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(user).RespJSON(c)
}

func (h userHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	//bind uri
	userUri := model.User{}
	errV, err := e.BindValidateURI(c, &userUri)
	if len(errV) > 0 {
		e.RespErrValidation(errV).RespJSON(c)
		return
	}

	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	err = h.userUc.DeleteUser(ctx, userUri.UserID)
	if err != nil {
		e.UnwrapErrorToResponse(err).RespJSON(c)
		return
	}

	res.RespSuccess(nil).RespJSON(c)
}
