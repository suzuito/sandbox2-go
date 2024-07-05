package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/environment"
	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/inject"
	admin_web "github.com/suzuito/sandbox2-go/photodx/service/admin/pkg/web"
	auth_web "github.com/suzuito/sandbox2-go/photodx/service/auth/pkg/web"
	authuser_web "github.com/suzuito/sandbox2-go/photodx/service/authuser/pkg/web"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
	user_web "github.com/suzuito/sandbox2-go/photodx/service/user/pkg/web"
)

func main() {
	ctx := context.Background()
	if err := setUp(ctx); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func setUp(
	ctx context.Context,
) error {
	env := environment.Environment{}
	if err := envconfig.Process("", &env); err != nil {
		return terrors.Wrapf("cannot load environment variables : %w", err)
	}
	slogHandler := inject.NewSlogHandler(&env)
	handler := clog.CustomHandler{
		Handler: slogHandler,
	}
	logger := slog.New(&handler)
	logic, err := inject.NewBusinessLogic(&env, logger)
	if err != nil {
		return terrors.Wrapf("NewBusinessLogic is failed : %w", err)
	}
	engine := gin.Default()
	common_web.SetRouter(
		engine,
		logger,
		env.CorsAllowOrigins,
		env.CorsAllowMethods,
		env.CorsAllowHeaders,
		env.CorsExposeHeaders,
	)
	auth_web.SetRouter(engine, logger, logic)
	admin_web.SetRouter(engine, logger, logic)
	user_web.SetRouter(engine, logger, logic)
	authuser_web.SetRouter(engine, logger, logic)
	if err := engine.Run(fmt.Sprintf(":%d", env.Port)); err != nil {
		return terrors.Wrapf("cannot run server : %w", err)
	}
	return nil
}
