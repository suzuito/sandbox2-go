package inject

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/go-sql-driver/mysql"
	"github.com/suzuito/sandbox2-go/blog2/internal/environment"
	"github.com/suzuito/sandbox2-go/blog2/internal/infra"
	"github.com/suzuito/sandbox2-go/blog2/internal/markdown2html"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func newUsecaseImplLocal(
	ctx context.Context,
	env *environment.Environment,
	arg *argNewUsecaseImpl,
) (
	usecase.Usecase,
	*slog.Logger,
	error,
) {
	slogHandler := clog.CustomHandler{
		Handler: newSlogHandlerText(slog.LevelDebug),
	}
	logger := slog.New(&slogHandler)
	mysqlConfig := mysql.Config{
		DBName:    "blog2",
		User:      env.DBUser,
		Net:       "tcp",
		Addr:      "127.0.0.1:3307",
		ParseTime: true,
	}
	pool, err := sql.Open(
		"mysql",
		mysqlConfig.FormatDSN(),
	)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}

	u := usecase.Impl{
		RepositoryArticle: &infra.RepositoryArticle{
			Pool: pool,
		},
		StorageArticle: &infra.StorageArticle{
			Cli:    arg.StorageClient,
			Bucket: env.ArticleMarkdownBucket,
		},
		StorageArticleFileDirectlyUploaded: &infra.StorageArticleFileDirectlyUploaded{
			Cli:    arg.StorageClient,
			Bucket: env.ArticleFileDirectlyUploadedBucket,
		},
		Markdown2HTML: &markdown2html.Markdown2HTMLImpl{},
		L:             logger,
	}
	return &u, logger, nil
}
