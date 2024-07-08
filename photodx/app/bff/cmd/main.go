package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
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
	resource, err := inject.NewResource(ctx, &env)
	if err != nil {
		return terrors.Wrapf("NewResource is failed : %w", err)
	}
	defer resource.Close()
	engine := gin.Default()
	common_web.Main(
		engine,
		resource.Logger,
		env.CorsAllowOrigins,
		env.CorsAllowMethods,
		env.CorsAllowHeaders,
		env.CorsExposeHeaders,
	)
	if err := admin_web.Main(
		engine,
		resource.Logger,
		resource.AdminAccessTokenJWTVerifier,
	); err != nil {
		return terrors.Wrapf("Main is failed : %w", err)
	}
	if err := auth_web.Main(
		engine,
		resource.Logger,
		resource.GormDB,
		resource.AdminRefreshTokenJWTCreator,
		resource.AdminRefreshTokenJWTVerifier,
		resource.AdminAccessTokenJWTCreator,
		resource.AdminAccessTokenJWTVerifier,
	); err != nil {
		return terrors.Wrapf("Main is failed : %w", err)
	}
	if err := user_web.Main(
		engine,
		resource.Logger,
		resource.UserAccessTokenJWTVerifier,
	); err != nil {
		return terrors.Wrapf("Main is failed : %w", err)
	}
	if err := authuser_web.Main(
		engine,
		resource.Logger,
		resource.GormDB,
		resource.UserRefreshTokenJWTCreator,
		resource.UserRefreshTokenJWTVerifier,
		resource.UserAccessTokenJWTCreator,
		resource.UserAccessTokenJWTVerifier,
	); err != nil {
		return terrors.Wrapf("Main is failed : %w", err)
	}
	if err := engine.Run(fmt.Sprintf(":%d", env.Port)); err != nil {
		return terrors.Wrapf("cannot run server : %w", err)
	}
	return nil
}
