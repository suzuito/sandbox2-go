package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func (t *UsecaseImpl) StartPipelinePeriodically(
	ctx context.Context,
) error {
	t.L.InfoContext(ctx, "StartPipelinePeriodically")
	crawlers, err := t.CrawlerRepository.GetCrawlerDefinitions(
		ctx,
		// goblog.CrawlerID,
		// goconnpass.CrawlerID,
		// golangweekly.CrawlerID,
		// knowledgeworkblogs.CrawlerID,
	)
	if err != nil {
		return terrors.Wrap(err)
	}
	for _, crw := range crawlers {
		loggerPerCrawler := t.L.With("crawlerID", crw.ID)
		loggerPerCrawler.InfoContext(ctx, "PublishCrawlEvent")
		if err := t.TriggerCrawlerQueue.PublishCrawlEvent(ctx, crw.ID, crawler.CrawlerInputData{}); err != nil {
			loggerPerCrawler.ErrorContext(ctx, "Failed to PublishCrawlEvent", "err", err)
			continue
		}
	}
	return nil
}
