package inject

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/blog2/internal/environment"
	"github.com/suzuito/sandbox2-go/blog2/internal/infra"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase"
	"github.com/suzuito/sandbox2-go/blog2/internal/web"
)

func NewImpl(ctx context.Context) (
	usecase.Usecase,
	*web.Impl,
	error,
) {
	env := environment.Environment{}
	if err := envconfig.Process("", &env); err != nil {
		return nil, nil, err
	}
	u := usecase.Impl{
		RepositoryArticle:      &infra.RepositoryArticle{},
		RepositoryArticleIndex: &infra.RepositoryArticleIndex{},
	}
	w := web.NewImpl(&u, &env)
	return &u, w, nil
}
