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
	logger := t.L.With("crawlerStarterID", crawlerStarterSettingID)
	logger.InfoContext(ctx, "StartPipelinePeriodically")
	settings, err := t.CrawlerRepository.GetCrawlerStarterSettings(ctx, crawlerStarterSettingID)
	if err != nil {
		logger.ErrorContext(ctx, "Failed to GetCrawlerStarterSettings", "err", err)
		return terrors.Wrap(err)
	}
	for _, setting := range settings {
		loggerPerCrawler := logger.With("crawlerID", setting.CrawlerID)
		loggerPerCrawler.InfoContext(ctx, "PublishDispatchCrawlEvent")
		if err := t.TriggerCrawlerQueue.PublishDispatchCrawlEvent(ctx, setting.CrawlerID, setting.CrawlerInputData); err != nil {
			loggerPerCrawler.ErrorContext(ctx, "Failed to PublishDispatchCrawlEvent", "err", err)
			continue
		}
	}
	return nil
}
