package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"shop_dev/controller"
)

type RegisterRoutersIn struct {
	fx.In
	Engine         *gin.Engine
	UserController *controller.UserController
}

func NewRegisterRouters(p RegisterRoutersIn) {
	r := p.Engine
	v1 := r.Group("/shop-dev/v1/")
	{
		auth := v1.Group("auth")
		{
			auth.POST("/signup", p.UserController.CreateNewUser)
			auth.POST("/login", p.UserController.LoginUser)
		}
	}
}
