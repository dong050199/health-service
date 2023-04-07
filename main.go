package main

import (
	"context"
	"errors"
	"fmt"
	"health_service/di"
	"health_service/pkg/config"
	"health_service/pkg/errormap"
	"health_service/pkg/infra"
	"health_service/pkg/middleware"
	"health_service/pkg/swagger"
	"health_service/service/http/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func registerSwaggerHandler(g *gin.Engine) {
	swaggerAPI := g.Group("/swagger")
	swag := swagger.NewSwagger()
	swaggerAPI.Use(swag.SwaggerHandler(false))
	swag.Register(swaggerAPI)
}

func startServer(ginEngine *gin.Engine, lifecycle fx.Lifecycle) {
	port := viper.GetString("PORT")
	server := http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}
	ginEngine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("run on port:", port)
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					fmt.Errorf("failed to listen and serve from server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func registerService(ginEngine *gin.Engine,
	userRouter router.UserRouter,
) {
	ginEngine.Use(middleware.Recover())
	gGroup := ginEngine.Group("api/v1")
	userRouter.Register(gGroup)
}

func main() {
	fx.New(
		fx.Invoke(errormap.Initialize),
		fx.Invoke(config.InitConfig),
		fx.Invoke(config.SetConfig),
		fx.Invoke(infra.InitMySQL),
		di.Module,
		fx.Provide(gin.Default),
		fx.Invoke(
			registerSwaggerHandler,
			registerService,
			startServer,
		),
	).Run()
}
