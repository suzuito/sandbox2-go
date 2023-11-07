package usecase

import (
	"bytes"
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
)

func (t *UsecaseImpl) Crawl2(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
	crawlerInputData crawler.CrawlerInputData,
) error {
	t.L.Infof(ctx, "Crawl %s", crawlerID)
	crawler, err := t.CrawlerFactory2.GetCrawler(ctx, crawlerID)
	if err != nil {
		t.L.Errorf(ctx, "Failed to GetCrawler : %+v", err)
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
	return nil
}
