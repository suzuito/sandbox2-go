package inject

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/environment"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase"
	"github.com/suzuito/sandbox2-go/photodx/internal/web"
	"github.com/suzuito/sandbox2-go/photodx/internal/web/presenter"
)

func NewWebImpl(
	env *environment.Environment,
	logger *slog.Logger,
	u usecase.Usecase,
) (*web.Impl, error) {
	return &web.Impl{
		U:                 u,
		P:                 &presenter.Impl{},
		L:                 logger,
		CorsAllowOrigins:  env.CorsAllowOrigins,
		CorsAllowMethods:  env.CorsAllowMethods,
		CorsAllowHeaders:  env.CorsAllowHeaders,
		CorsExposeHeaders: env.CorsExposeHeaders,
	}, nil
}
