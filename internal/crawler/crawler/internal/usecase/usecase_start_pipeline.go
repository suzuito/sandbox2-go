package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/crawler/goblog"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/crawler/goconnpass"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/crawler/golangweekly"
)

func (t *UsecaseImpl) StartPipelinePeriodically(
	ctx context.Context,
) error {
	t.L.Infof(ctx, "StartPipelinePeriodically")
	crawlers := t.CrawlerFactory.GetCrawlers(
		ctx,
		goblog.CrawlerID,
		goconnpass.CrawlerID,
		golangweekly.CrawlerID,
	)
	for _, crawler := range crawlers {
		t.L.Infof(ctx, "Start %s (%s)", crawler.ID(), crawler.Name())
		if err := t.Queue.PublishCrawlEvent(ctx, crawler.ID()); err != nil {
			t.L.Errorf(ctx, "PublishCrawlEvent of '%s' is failed : %+v", crawler.ID(), err)
			continue
		}
	}
	return nil
}
