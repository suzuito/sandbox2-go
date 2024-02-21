package inject

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/blog2/internal/environment"
	"github.com/suzuito/sandbox2-go/blog2/internal/infra"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase"
	"github.com/suzuito/sandbox2-go/blog2/internal/web"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
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
	slogHandler := clog.CustomHandler{
		Handler: slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			},
		),
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
		// fmt.Sprintf(
		// 	"%s@tcp(127.0.0.1:3307)/blog2?charset=utf8mb4&parseTime=True",
		// 	env.DBUser,
		// ),
		mysqlConfig.FormatDSN(),
	)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	u := usecase.Impl{
		RepositoryArticle: &infra.RepositoryArticle{
			Pool: pool,
		},
		RepositoryArticleIndex: &infra.RepositoryArticleIndex{},
		L:                      logger,
	}
	w := web.Impl{
		U:          &u,
		P:          web.NewPresenter(),
		L:          logger,
		AdminToken: env.AdminToken,
	}
	return &u, &w, nil
}
