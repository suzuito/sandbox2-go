package inject

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/storage"
	"github.com/google/go-github/v50/github"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/internal/blog/infra/bgcp"
	"github.com/suzuito/sandbox2-go/internal/blog/infra/bgithub"
	"github.com/suzuito/sandbox2-go/internal/blog/infra/bmysql"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase/markdown2html"
	"github.com/suzuito/sandbox2-go/internal/blog/web"
	"golang.org/x/oauth2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Environment struct {
	DirPathArticleSources       string `required:"false" split_words:"true"`
	DirPathArticleHTML          string `required:"false" split_words:"true"`
	DBUser                      string `required:"false" split_words:"true"`
	DBPassword                  string `required:"false" split_words:"true"`
	DBGCPRegion                 string `required:"false" envconfig:"DB_GCP_REGION"`
	ArticleSourcesGHOwner       string `required:"false" envconfig:"ARTICLE_SOURCES_GH_OWNER"`
	ArticleSourcesGHRepo        string `required:"false" envconfig:"ARTICLE_SOURCES_GH_REPO"`
	ArticleSourcesGHAccessToken string `required:"false" envconfig:"ARTICLE_SOURCES_GH_ACCESS_TOKEN"`
	ArticleSourcesDirPath       string `required:"false" split_words:"true"`
	ArticleHTMLBucket           string `required:"false" envconfig:"ARTICLE_HTML_BUCKET"`
	SiteOrigin                  string `required:"false" split_words:"true"`
	Env                         string `required:"true" split_words:"true"`
	AdminPassword               string `required:"true" split_words:"true"`
}

func NewUsecaseImpl(ctx context.Context) (*usecase.UsecaseImpl, *Environment, *web.WebSetting, error) {
	var env Environment
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, nil, nil, err
	}

	db, err := newGormDB(&env)
	if err != nil {
		return nil, nil, nil, err
	}

	cliGCS, err := newGoogleCloudStorageClient(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	u := usecase.UsecaseImpl{
		RepositoryArticleSource: newRepositoryArticleSource(ctx, &env),
		RepositoryArticleHTML: &bgcp.RepositoryArticleHTML{
			Cli:    cliGCS,
			Bucket: env.ArticleHTMLBucket,
		},
		RepositoryArticle: &bmysql.RepositoryArticle{
			DB: db,
		},
		Markdown2HTML: &markdown2html.Markdown2HTMLImpl{},
	}
	s := web.WebSetting{
		SiteOrigin:       env.SiteOrigin,
		DirPathTemplates: "internal/blog/web/templates",
		DirPathCSS:       "internal/blog/web/css",
		DirPathImages:    "internal/blog/web/images",
		AdminPassword:    env.AdminPassword,
	}
	if env.Env == "godzilla" {
		s.NoIndex = false
	} else {
		s.NoIndex = true
	}
	return &u, &env, &s, nil
}

func newRepositoryArticleSource(
	ctx context.Context,
	env *Environment,
) usecase.RepositoryArticleSource {
	httpCli := oauth2.NewClient(
		ctx,
		oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: env.ArticleSourcesGHAccessToken,
			},
		),
	)
	cli := github.NewClient(httpCli)
	return &bgithub.RepositoryArticleSource{
		Client:  cli,
		Owner:   env.ArticleSourcesGHOwner,
		Repo:    env.ArticleSourcesGHRepo,
		BaseDir: env.ArticleSourcesDirPath,
	}
}

func newGormDB(
	env *Environment,
) (*gorm.DB, error) {
	dst := ""
	var dbLogger logger.Interface
	if env.Env == "dev" {
		dst = "root:@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True"
		dbLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Duration(0), // Slow SQL threshold
				LogLevel:                  logger.Info,      // Log level
				IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,             // Disable color
			},
		)
	} else {
		gcpProjectID, err := metadata.ProjectID()
		if err != nil {
			return nil, fmt.Errorf("metadata.ProjectID is failed: %w", err)
		}
		fmt.Printf("hoge %s\n", gcpProjectID)
		dst = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s:%s:sandbox-instance)/blog?charset=utf8mb4&parseTime=True",
			env.DBUser,
			env.DBPassword,
			gcpProjectID,
			env.DBGCPRegion,
		)
		dbLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Duration(10) * time.Second, // Slow SQL threshold
				LogLevel:                  logger.Warn,                     // Log level
				IgnoreRecordNotFoundError: true,                            // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,                           // Disable color
			},
		)
	}
	db, err := openMysql(dst, dbLogger)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect database '%s' : %w", dst, err)
	}

	return db, nil
}

func openMysql(
	dst string,
	dbLogger logger.Interface,
) (*gorm.DB, error) {
	config := gorm.Config{
		Logger: dbLogger,
	}
	db, err := gorm.Open(
		mysql.Open(dst),
		&config,
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newGoogleCloudStorageClient(
	ctx context.Context,
) (*storage.Client, error) {
	return storage.NewClient(ctx)
}
