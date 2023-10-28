package usecase

import (
	"bytes"
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/entity/crawler"
)

func (t *UsecaseImpl) CrawlOnGCF(
	ctx context.Context,
	rawBytes []byte,
) error {
	crawlerID, err := t.Queue.RecieveCrawlEvent(ctx, rawBytes)
	if err != nil {
		t.L.Errorf(ctx, "Failed to RecieveCrawlEvent : %+v", err)
		return terrors.Wrap(err)
	}
	ctx = context.WithValue(ctx, "crawlerId", crawlerID)
	if err := t.Crawl(ctx, crawlerID); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *UsecaseImpl) Crawl(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
) error {
	t.L.Infof(ctx, "Crawl %s", crawlerID)
	crawler, err := t.CrawlerFactory.GetCrawler(ctx, crawlerID)
	if err != nil {
		t.L.Errorf(ctx, "Failed to GetCrawler : %+v", err)
		return terrors.Wrap(err)
	}
	fetcher, err := crawler.NewFetcher(ctx)
	if err != nil {
		t.L.Errorf(ctx, "Failed to NewFetcher : %+v", err)
		return terrors.Wrap(err)
	}
	data := []byte{}
	w := bytes.NewBuffer(data)
	if err := fetcher.Fetch(ctx, w); err != nil {
		t.L.Errorf(ctx, "Failed to Fetch : %+v", err)
		return terrors.Wrap(err)
	}
	parser, err := crawler.NewParser(ctx)
	if err != nil {
		t.L.Errorf(ctx, "Failed to NewParser : %+v", err)
		return terrors.Wrap(err)
	}
	timeSeriesData, err := parser.Parse(ctx, w)
	if err != nil {
		t.L.Errorf(ctx, "Failed to Parse : %+v", err)
		return terrors.Wrap(err)
	}
	for _, data := range timeSeriesData {
		t.L.Debugf(ctx, "fetched data %s", string(data.GetID()))
	}
	publisher, err := crawler.NewPublisher(ctx)
	if err != nil {
		t.L.Errorf(ctx, "Failed to NewPublisher : %+v", err)
		return terrors.Wrap(err)
	}
	if err := publisher.Publish(ctx, timeSeriesData...); err != nil {
		t.L.Errorf(ctx, "Failed to Publish : %+v", err)
		return terrors.Wrap(err)
	}
	return nil
}
