package usecase

import (
	"bytes"
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func (t *UsecaseImpl) CrawlOnGCF(
	ctx context.Context,
	rawBytes []byte,
) error {
	crawlerID, crawlerInputData, err := t.TriggerCrawlerQueue.RecieveCrawlEvent(ctx, rawBytes)
	if err != nil {
		t.L.ErrorContext(ctx, "Failed to RecieveCrawlEvent", "err", err)
		return terrors.Wrap(err)
	}
	if err := t.Crawl(ctx, crawlerID, crawlerInputData); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *UsecaseImpl) Crawl(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
	crawlerInputData crawler.CrawlerInputData,
) error {
	logger := t.L.With("crawlerID", crawlerID)
	logger.InfoContext(ctx, "Crawl")
	crawlerDefinition, err := t.CrawlerRepository.GetCrawlerDefinition(ctx, crawlerID)
	if err != nil {
		logger.ErrorContext(ctx, "Failed to GetCrawler", "err", err)
		return terrors.Wrap(err)
	}
	crawler, err := t.CrawlerFactory.Get(ctx, crawlerDefinition)
	if err != nil {
		logger.ErrorContext(ctx, "Failed to CrawlerFactory.Get", "err", err)
		return terrors.Wrap(err)
	}
	data := []byte{}
	w := bytes.NewBuffer(data)
	if err := crawler.Fetcher.Do(ctx, logger, w, crawlerInputData); err != nil {
		t.L.ErrorContext(ctx, "Failed to Fetch", "err", err)
		return terrors.Wrap(err)
	}
	timeSeriesData, err := crawler.Parser.Do(ctx, w, crawlerInputData)
	if err != nil {
		t.L.ErrorContext(ctx, "Failed to Parse", "err", err)
		return terrors.Wrap(err)
	}
	for _, data := range timeSeriesData {
		t.L.DebugContext(ctx, "fetched data", "timeSeriesDataID", string(data.GetID()))
	}
	if err := crawler.Publisher.Do(ctx, crawlerInputData, timeSeriesData...); err != nil {
		t.L.ErrorContext(ctx, "Failed to Publish", "err", err)
		return terrors.Wrap(err)
	}
	return nil
}
