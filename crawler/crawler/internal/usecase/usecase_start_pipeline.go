package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/goblog"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/goconnpass"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/golangweekly"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/knowledgeworkblogs"
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
		knowledgeworkblogs.CrawlerID,
	)
	for _, crw := range crawlers {
		t.L.Infof(ctx, "Start %s (%s)", crw.ID(), crw.Name())
		if err := t.Queue.PublishCrawlEvent(ctx, crw.ID(), crawler.CrawlerInputData{}); err != nil {
			t.L.Errorf(ctx, "PublishCrawlEvent of '%s' is failed : %+v", crw.ID(), err)
			continue
		}
	}
	return nil
}
