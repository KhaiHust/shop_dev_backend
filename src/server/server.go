package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"shop_dev/config"
)

func NewGinEngine(lc fx.Lifecycle, config *config.Config) *gin.Engine {
	engine := gin.New()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting HTTP server")

			go func() {
				addr := fmt.Sprintf("%s:%s", config.HttpConfig.Host, config.HttpConfig.Port)
				if err := engine.Run(addr); err != nil {
					log.Fatal(fmt.Sprint("HTTP server failed to start %w", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping HTTP server")
			return nil
		},
	})
	return engine
}
