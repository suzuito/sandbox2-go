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
	"github.com/suzuito/sandbox2-go/photodx/internal/web"
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
	uc, err := inject.NewUsecase(&env, logger)
	if err != nil {
		return terrors.Wrapf("NewUsecase is failed : %w", err)
	}
	webImpl, err := inject.NewWebImpl(&env, logger, uc)
	if err != nil {
		return terrors.Wrapf("NewWebImpl is failed : %w", err)
	}
	engine := gin.Default()
	web.SetRouter(engine, webImpl)
	if err := engine.Run(fmt.Sprintf(":%d", env.Port)); err != nil {
		return terrors.Wrapf("cannot run server : %w", err)
	}
	return nil
}
