package inject

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-cz/devslog"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/app/bff/internal/environment"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

type Resource struct {
	GormDB *gorm.DB
	Logger *slog.Logger

	AdminAccessTokenJWTVerifier  auth.JWTVerifier
	AdminAccessTokenJWTCreator   auth.JWTCreator
	AdminRefreshTokenJWTVerifier auth.JWTVerifier
	AdminRefreshTokenJWTCreator  auth.JWTCreator
	UserAccessTokenJWTVerifier   auth.JWTVerifier
	UserAccessTokenJWTCreator    auth.JWTCreator
	UserRefreshTokenJWTVerifier  auth.JWTVerifier
	UserRefreshTokenJWTCreator   auth.JWTCreator
}

func (t *Resource) Close() {
	t.Logger.Info("Close resources")
}

func NewResource(
	ctx context.Context,
	env *environment.Environment,
) (*Resource, error) {
	resource := Resource{}
	switch env.Env {
	case "local":
		setLocalResource(env, &resource)
	default:
		return nil, terrors.Wrapf("undefined env resource : %s", env.Env)
	}
	setJWTResource(env, &resource)
	return &resource, nil
}

func setLocalResource(
	env *environment.Environment,
	resource *Resource,
) error {
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
	gormdb, err := gorm.Open(
		gorm_mysql.Open(mysqlConfig.FormatDSN()),
		&gorm.Config{
			Logger: gorm_logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				gorm_logger.Config{
					Colorful:      true,
					LogLevel:      gorm_logger.Info,
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

func setJWTResource(
	env *environment.Environment,
	resource *Resource,
) error {
	// AdminAccessToken
	adminRefreshTokenProcessor := auth.JWTHS256{
		PrivateKey: []byte(env.JWTAdminRefreshTokenSigningPrivateKey),
	}
	resource.AdminRefreshTokenJWTCreator = &adminRefreshTokenProcessor
	resource.AdminRefreshTokenJWTVerifier = &adminRefreshTokenProcessor
	adminAccessTokenJWTPrivateKeyBytes, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(env.JWTAdminAccessTokenSigningPrivateKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.AdminAccessTokenJWTCreator = &auth.JWTCreatorRS256{
		PrivateKey: adminAccessTokenJWTPrivateKeyBytes,
	}
	adminAccessTokenJWTPublicKeyBytes, err := jwt.ParseRSAPublicKeyFromPEM([]byte(env.JWTAdminAccessTokenSigningPublicKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.AdminAccessTokenJWTVerifier = &auth.JWTVerifiers{
		Verifiers: []auth.JWTVerifier{
			&auth.JWTVerifierRS256{
				PublicKey: adminAccessTokenJWTPublicKeyBytes,
			},
		},
	}
	// UserAccessToken
	userRefreshTokenProcessor := auth.JWTHS256{
		PrivateKey: []byte(env.JWTUserRefreshTokenSigningPrivateKey),
	}
	resource.UserRefreshTokenJWTCreator = &userRefreshTokenProcessor
	resource.UserRefreshTokenJWTVerifier = &userRefreshTokenProcessor
	userAccessTokenJWTPrivateKeyBytes, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(env.JWTUserAccessTokenSigningPrivateKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.UserAccessTokenJWTCreator = &auth.JWTCreatorRS256{
		PrivateKey: userAccessTokenJWTPrivateKeyBytes,
	}
	userAccessTokenJWTPublicKeyBytes, err := jwt.ParseRSAPublicKeyFromPEM([]byte(env.JWTUserAccessTokenSigningPublicKey))
	if err != nil {
		return terrors.Wrap(err)
	}
	resource.UserAccessTokenJWTVerifier = &auth.JWTVerifiers{
		Verifiers: []auth.JWTVerifier{
			&auth.JWTVerifierRS256{
				PublicKey: userAccessTokenJWTPublicKeyBytes,
			},
		},
	}
	return nil
}
