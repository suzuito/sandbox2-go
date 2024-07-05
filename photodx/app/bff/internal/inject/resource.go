package inject

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-cz/devslog"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/environment"
)

type Resource struct {
	Pool   *sql.DB
	Logger *slog.Logger
}

func (t *Resource) Close() {
	t.Logger.Info("Close resources")
	if err := t.Pool.Close(); err != nil {
		t.Logger.Error("Pool.Close() is failed", "err", err)
	}
}

func NewResource(
	ctx context.Context,
	env *environment.Environment,
) (*Resource, error) {
	switch env.Env {
	case "local":
		return newLocalResource(env)
	}
	return nil, terrors.Wrapf("undefined env resource : %w", env.Env)
}

func newLocalResource(
	env *environment.Environment,
) (*Resource, error) {
	resource := Resource{}
	// logger
	var level slog.Level
	if err := level.UnmarshalText([]byte(env.LogLevel)); err != nil {
		fmt.Printf("use LogLevel 'DEBUG' because cannot parse LOG_LEVEL : %s", env.LogLevel)
		level = slog.LevelDebug
	}
	slogHandler := devslog.NewHandler(os.Stdout, &devslog.Options{
		HandlerOptions: &slog.HandlerOptions{
			Level:     level,
			AddSource: true,
		},
	})
	slogCustomHandler := clog.CustomHandler{
		Handler: slogHandler,
	}
	resource.Logger = slog.New(&slogCustomHandler)
	// mysql
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
	resource.Pool = pool
	return &resource, nil
}
