package inject

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/environment"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

func setPrdResource(
	env *environment.Environment,
	resource *Resource,
) error {
	// logger
	var level slog.Level
	if err := level.UnmarshalText([]byte(env.LogLevel)); err != nil {
		fmt.Printf("use LogLevel 'DEBUG' because cannot parse LOG_LEVEL : %s", env.LogLevel)
		level = slog.LevelDebug
	}
	slogHandler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:     level,
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.LevelKey {
					a.Key = "severity"
				}
				return a
			},
		},
	)
	slogCustomHandler := clog.CustomHandler{
		Handler: slogHandler,
	}
	resource.Logger = slog.New(&slogCustomHandler)
	// mysql
	mysqlConfig := mysql.Config{
		DBName:               env.DBName,
		User:                 env.DBUser,
		Passwd:               env.DBPassword,
		Net:                  "unix",
		Addr:                 env.DBInstanceUnixSocket,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	gormdb, err := gorm.Open(
		gorm_mysql.Open(mysqlConfig.FormatDSN()),
		&gorm.Config{
			Logger: gorm_logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				gorm_logger.Config{
					Colorful:      false,
					LogLevel:      gorm_logger.Error,
					SlowThreshold: time.Second,
				},
			),
		},
	)
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.GormDB = gormdb
	return nil
}
