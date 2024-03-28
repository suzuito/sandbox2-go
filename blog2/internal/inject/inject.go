package inject

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/blog2/internal/environment"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase"
	"github.com/suzuito/sandbox2-go/blog2/internal/web"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func NewImpl(ctx context.Context) (
	usecase.Usecase,
	*web.Impl,
	error,
) {
	var err error
	arg := argNewImpl{}
	if err := envconfig.Process("", &arg.Env); err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	arg.StorageClient, err = storage.NewClient(ctx)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	if arg.Env.Env == "dev" {
		return newImplLocal(ctx, &arg)
	}
	return newImpl(ctx, &arg)
}

type argNewImpl struct {
	Env           environment.Environment
	StorageClient *storage.Client
}
