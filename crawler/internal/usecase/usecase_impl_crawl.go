package usecase

import (
	"bytes"
	"context"
	"fmt"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func (t *UsecaseImpl) CrawlOnGCF(
	ctx context.Context,
	rawBytes []byte,
) error {
	crawlerID, crawlerInputData, err := t.TriggerCrawlerQueue.RecieveCrawlEvent(ctx, rawBytes)
	if err != nil {
		t.L.Errorf(ctx, "Failed to RecieveCrawlEvent : %+v", err)
		return terrors.Wrap(err)
	}
	ctx = context.WithValue(ctx, "crawlerId", crawlerID) // WARN: should not use built-in type string as key for value; define your own type to avoid collisions (SA1029)go-staticcheck
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
	t.L.Infof(ctx, "Crawl %s", crawlerID)
	crawlerDefinition, err := t.CrawlerRepository.GetCrawlerDefinition(ctx, crawlerID)
	if err != nil {
		t.L.Errorf(ctx, "Failed to GetCrawler : %+v", err)
		return terrors.Wrap(err)
	}
	crawler, err := t.CrawlerFactory.Get(ctx, crawlerDefinition)
	if err != nil {
		t.L.Errorf(ctx, "Failed to CrawlerFactory.Get : %+v", err)
		return terrors.Wrap(err)
	}
	data := []byte{}
	w := bytes.NewBuffer(data)
	if err := crawler.Fetcher.Do(ctx, w, crawlerInputData); err != nil {
		t.L.Errorf(ctx, "Failed to Fetch : %+v", err)
		return terrors.Wrap(err)
	}
	timeSeriesData, err := crawler.Parser.Do(ctx, w, crawlerInputData)
	if err != nil {
		t.L.Errorf(ctx, "Failed to Parse : %+v", err)
		return terrors.Wrap(err)
	}
	for _, data := range timeSeriesData {
		t.L.Debugf(ctx, "fetched data %s", string(data.GetID()))
	}
	if err := crawler.Publisher.Do(ctx, crawlerInputData, timeSeriesData...); err != nil {
		t.L.Errorf(ctx, "Failed to Publish : %+v", err)
		return terrors.Wrap(err)
	}
	return fmt.Errorf("not impl")
}
