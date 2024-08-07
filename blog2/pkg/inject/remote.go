package inject

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/go-sql-driver/mysql"
	"github.com/suzuito/sandbox2-go/blog2/internal/infra"
	"github.com/suzuito/sandbox2-go/blog2/internal/markdown2html"
	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile"
	internal_usecase "github.com/suzuito/sandbox2-go/blog2/internal/usecase"
	"github.com/suzuito/sandbox2-go/blog2/pkg/environment"
	"github.com/suzuito/sandbox2-go/blog2/pkg/usecase"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func newUsecaseImpl(
	ctx context.Context,
	env *environment.Environment,
	arg *argNewUsecaseImpl,
) (
	usecase.Usecase,
	*slog.Logger,
	error,
) {
	slogHandler := clog.CustomHandler{
		Handler: newSlogHandlerJSON(slog.LevelDebug),
	}
	logger := slog.New(&slogHandler)
	mysqlConfig := mysql.Config{
		DBName: env.DBName,
		User:   env.DBUser,
		Passwd: env.DBPassword,
		Net:    "unix",
		Addr:   env.DBInstanceUnixSocket,
		// Addr:                 fmt.Sprintf("/cloudsql/%s:asia-northeast1:sandbox-instance", gcpProjectID),
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	pool, err := sql.Open(
		"mysql",
		mysqlConfig.FormatDSN(),
	)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}

	u := internal_usecase.NewImpl(
		&infra.RepositoryArticle{
			Pool: pool,
		},
		&infra.StorageArticle{
			Cli:    arg.StorageClient,
			Bucket: env.ArticleMarkdownBucket,
		},
		&infra.StorageFile{
			Cli:    arg.StorageClient,
			Bucket: env.FileBucket,
		},
		&infra.StorageFileThumbnail{
			Cli:    arg.StorageClient,
			Bucket: env.FileThumbnailBucket,
		},
		articlefile.NewImageConverter(),
		&markdown2html.Markdown2HTMLImpl{},
		articlefile.NewFileTypeDetector(),
		logger,
	)
	return u, logger, nil
}
