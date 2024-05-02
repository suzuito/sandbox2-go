package inject

import (
	"context"
	"log/slog"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/web"
	"github.com/suzuito/sandbox2-go/blog2/pkg/environment"
	"github.com/suzuito/sandbox2-go/blog2/pkg/usecase"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func NewUsecaseImpl(ctx context.Context, env *environment.Environment) (
	usecase.Usecase,
	*slog.Logger,
	error,
) {
	var err error
	arg := argNewUsecaseImpl{}
	arg.StorageClient, err = storage.NewClient(ctx)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	var u usecase.Usecase
	var logger *slog.Logger
	if env.Env == "dev" {
		u, logger, err = newUsecaseImplLocal(ctx, env, &arg)
	} else {
		u, logger, err = newUsecaseImpl(ctx, env, &arg)
	}
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	return u, logger, nil
}

type argNewUsecaseImpl struct {
	StorageClient *storage.Client
}

func NewWebImpl(
	ctx context.Context,
	env *environment.Environment,
	u usecase.Usecase,
	logger *slog.Logger,
) *web.Impl {
	w := web.Impl{
		U:                    u,
		P:                    web.NewPresenter(),
		L:                    logger,
		AdminToken:           env.AdminToken,
		BaseURLFile:          env.BaseURLFile,
		BaseURLFileThumbnail: env.BaseURLFileThumbnail,
	}
	return &w
}
