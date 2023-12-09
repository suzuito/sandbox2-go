package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func (t *UsecaseImpl) DispatchCrawlOnGCF(
	ctx context.Context,
	rawBytes []byte,
) error {
	crawlerID, crawlerInputData, err := t.TriggerCrawlerQueue.RecieveCrawlEvent(ctx, rawBytes)
	if err != nil {
		t.L.ErrorContext(ctx, "Failed to RecieveCrawlEvent", "err", err)
		return terrors.Wrap(err)
	}
	ctx = context.WithValue(ctx, "crawlerId", crawlerID) // WARN: should not use built-in type string as key for value; define your own type to avoid collisions (SA1029)go-staticcheck
	if err := t.DispatchCrawl(ctx, crawlerID, crawlerInputData); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *UsecaseImpl) DispatchCrawl(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
	crawlerInputData crawler.CrawlerInputData,
) error {
	logger := t.L.With("crawlerID", crawlerID)
	logger.InfoContext(ctx, "DispatchCrawl")
	setting, err := t.CrawlerConfigurationRepository.GetDispatchCrawlSetting(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "Failed to GetDispatchCrawlSetting", "err", err)
		return terrors.Wrap(err)
	}
	crawlFunctionID, exists := setting.CrawlFunctionIDMapping[crawlerID]
	if !exists {
		crawlFunctionID = setting.DefaultCrawlFunctionID
	}
	logger.InfoContext(ctx, "PublishCrawlEvent", "crawlFunctionID", crawlFunctionID)
	if err := t.TriggerCrawlerQueue.PublishCrawlEvent(ctx, crawlerID, crawlerInputData, crawlFunctionID); err != nil {
		logger.ErrorContext(ctx, "Failed to PublishCrawlEvent", "err", err)
		return terrors.Wrap(err)
	}
	return nil
}
