package inject

import (
	"context"
	"database/sql"
	"log/slog"

	"cloud.google.com/go/firestore"
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
	firestoreClient, err := firestore.NewClient(ctx, "suzuito-minilla") // CloudFunctionとの連携の必要性から、ローカルではなくminillaを使う
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	mysqlConfig := mysql.Config{
		DBName:    env.DBName,
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

	u := internal_usecase.NewImpl(
		&infra.RepositoryArticle{
			Pool: pool,
		},
		&infra.StorageArticle{
			Cli:    arg.StorageClient,
			Bucket: env.ArticleMarkdownBucket,
		},
		&infra.StorageFileUploaded{
			Cli:    arg.StorageClient,
			Bucket: env.FileUploadedBucket,
		},
		&infra.StorageFile{
			Cli:    arg.StorageClient,
			Bucket: env.FileBucket,
		},
		&infra.StorageFileThumbnail{
			Cli:    arg.StorageClient,
			Bucket: env.FileThumbnailBucket,
		},
		&infra.RepositoryFileUploaded{
			Cli: firestoreClient,
		},
		articlefile.NewImageConverter(),
		&markdown2html.Markdown2HTMLImpl{},
		articlefile.NewFileTypeDetector(),
		logger,
	)
	return u, logger, nil
}
