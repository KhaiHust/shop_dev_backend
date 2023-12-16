package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"shop_dev/config"
	"shop_dev/controller"
	"shop_dev/db"
	"shop_dev/repository"
	"shop_dev/router"
	"shop_dev/server"
	"shop_dev/service"
)

func All() fx.Option {
	return fx.Options(
		//init
		fx.Provide(server.NewGinEngine),
		fx.Invoke(config.NewLogConfig),
		fx.Invoke(func(*gin.Engine) {}),
		fx.Provide(config.NewConfig),
		fx.Invoke(router.NewRegisterRouters),
		fx.Provide(db.NewInitDatabase),

		//inject repository
		fx.Provide(repository.NewUserRepository),

		//inject service
		fx.Provide(service.NewDatabaseTransaction),
		fx.Provide(service.
			NewUserService),

		//inject controller
		fx.Provide(controller.NewUserController),
	)
}
