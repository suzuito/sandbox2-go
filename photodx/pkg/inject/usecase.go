package inject

import (
	"database/sql"
	"log/slog"

	"github.com/go-sql-driver/mysql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/infra/rdb"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service"
	"github.com/suzuito/sandbox2-go/photodx/pkg/environment"
)

func NewUsecase(
	env *environment.Environment,
	logger *slog.Logger,
) (usecase.Usecase, error) {
	mysqlConfig := mysql.Config{
		DBName:               env.DBName,
		User:                 env.DBUser,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3308",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	pool, err := sql.Open(
		"mysql",
		mysqlConfig.FormatDSN(),
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	repository := rdb.Repository{
		Pool: pool,
	}
	s := service.Impl{
		Repository: repository,
	}
	u := usecase.Impl{
		S: &s,
	}
	return &u, nil
}
