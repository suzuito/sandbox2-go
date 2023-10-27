package main

import (
	"context"
	"os"

	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/crawler/goconnpass"
)

func main() {
	ctx := context.Background()
	u, err := NewUsecaseLocal(ctx)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
	crawlerID := goconnpass.CrawlerID
	// crawlerID = golangweekly.CrawlerID
	if err := u.Crawl(ctx, crawlerID); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
}
