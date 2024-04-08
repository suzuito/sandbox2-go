package inject

import (
	"context"
	"database/sql"
	"log/slog"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
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

	pubsubClient, err := pubsub.NewClient(ctx, "suzuito-minilla") // CloudFunctionとの連携の必要性から、ローカルではなくminillaを使う
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}

	u := internal_usecase.Impl{
		RepositoryArticle: &infra.RepositoryArticle{
			Pool: pool,
		},
		StorageArticle: &infra.StorageArticle{
			Cli:    arg.StorageClient,
			Bucket: env.ArticleMarkdownBucket,
		},
		StorageFileUploaded: &infra.StorageFileUploaded{
			Cli:    arg.StorageClient,
			Bucket: env.ArticleFileUploadedBucket,
		},
		StorageFile: &infra.StorageFile{
			Cli:    arg.StorageClient,
			Bucket: env.ArticleFileBucket,
		},
		RepositoryFileUploaded: &infra.RepositoryFileUploaded{
			Cli: firestoreClient,
		},
		FunctionTriggerStartFileUploadedProcess: &infra.FunctionTrigger{
			Cli:     pubsubClient,
			TopicID: env.FunctionTriggerTopicIDStartFileUploadedProcess,
		},
		FileImageConverter: articlefile.NewImageConverter(),
		Markdown2HTML:      &markdown2html.Markdown2HTMLImpl{},
		L:                  logger,
	}
	return &u, logger, nil
}
