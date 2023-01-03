package module

import (
	h "github.com/deall-users/internal/module/handler"
	rp "github.com/deall-users/internal/module/repository"
	uc "github.com/deall-users/internal/module/usecase"

	"github.com/deall-users/pkg/gorm"
)

type HTTPHandler struct {
	h.MiddlewareHandler
	h.UserHandler
	h.AuthHandler
}

type Usecase struct {
	uc.UserUsecase
	uc.AuthUsecase
}

type Repository struct {
	rp.UserRepository
}

func NewRepository(db *gorm.DbPgSql) *Repository {
	userRp := rp.NewUserRepo(db)
	return &Repository{userRp}
}

func NewUsecase(repo *Repository) *Usecase {
	userUc := uc.NewUserUsecase(repo.UserRepository)
	authUc := uc.NewAuthUsecase(repo.UserRepository)

	return &Usecase{
		userUc,
		authUc,
	}
}

func NewHTTPHandler(ucase *Usecase) *HTTPHandler {
	middleware := h.NewMiddlewareHandler(ucase.UserUsecase)
	userHand := h.NewUserHandler(ucase.UserUsecase)
	authHand := h.NewAuthHandler(ucase.AuthUsecase)

	return &HTTPHandler{
		middleware,
		userHand,
		authHand,
	}
}
