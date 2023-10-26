package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/cmd/internal/inject"
	"github.com/suzuito/sandbox2-go/internal/blog/web"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
)

func main() {
	clog.L.AddKey("error", "code")
	ctx := context.Background()
	u, _, s, err := inject.NewUsecaseImpl(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	ctrl := web.ControllerImpl{
		UC:                      u,
		RepositoryArticle:       u.RepositoryArticle,
		RepositoryArticleSource: u.RepositoryArticleSource,
		RepositoryArticleHTML:   u.RepositoryArticleHTML,
		Markdown2HTML:           u.Markdown2HTML,
		Presenters:              web.NewPresenters(),
		Setting:                 s,
	}
	e := gin.New()
	e.Use(gin.Recovery())
	web.SetRouter(e, &ctrl)
	e.Run()
}
