package gin

import (
	"github.com/deall-users/internal/module"
	"github.com/gin-gonic/gin"
)

type HTTPRouter struct {
	app *gin.Engine
	*module.HTTPHandler
}

func NewHTTPRouter(app *gin.Engine, uc *module.Usecase) *HTTPRouter {

	return &HTTPRouter{
		app,
		module.NewHTTPHandler(uc),
	}
}

func (r *HTTPRouter) InitRouters() {
	router := r.app.Group("/api/v1")
	router.Use(r.SetCors())
	router.OPTIONS("/*path", r.SetCors())

	{
		router.GET("/encrypt", r.Encryption)
		router.GET("/decrypt", r.Decryption)
	}

	{
		router.POST("/login", r.Login)
		router.GET("/logout", r.TokenValidation(false), r.Logout)
		router.GET("/token-validation", r.TokenValidation(false), r.UserTokenValidation)
		router.GET("/token-refresh", r.TokenValidation(true), r.UserRefreshToken)
	}

	// users
	userRoute := router.Group("/user")
	{
		userRoute.Use(r.TokenValidation(false), r.Permission())
		userRoute.POST("", r.CreateUser)
		userRoute.PUT("", r.UpdateUser)
		userRoute.GET("", r.FindAllUser)
		userRoute.GET("/:user_id", r.FindUserByID)
		userRoute.DELETE("/:user_id", r.DeleteUser)
	}

}
