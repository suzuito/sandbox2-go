package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func (t *UsecaseImpl) StartPipelinePeriodically(
	ctx context.Context,
	crawlerStarterSettingID crawler.CrawlerStarterSettingID,
) error {
	t.L.InfoContext(ctx, "StartPipelinePeriodically")
	settings, err := t.CrawlerRepository.GetCrawlerStarterSettings(ctx, crawlerStarterSettingID)
	if err != nil {
		return terrors.Wrap(err)
	}
	for _, setting := range settings {
		loggerPerCrawler := t.L.With("crawlerID", setting.CrawlerID)
		loggerPerCrawler.InfoContext(ctx, "PublishCrawlEvent")
		if err := t.TriggerCrawlerQueue.PublishCrawlEvent(ctx, setting.CrawlerID, setting.CrawlerInputData); err != nil {
			loggerPerCrawler.ErrorContext(ctx, "Failed to PublishCrawlEvent", "err", err)
			continue
		}
	}
	return nil
}
