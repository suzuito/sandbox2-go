package crawlerfactory

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
)

type CrawlerFactory2Impl struct {
	crawlers map[crawler.CrawlerID]crawler.Crawler2
}

func (t *CrawlerFactory2Impl) GetCrawler(ctx context.Context, crawlerID crawler.CrawlerID) (*crawler.Crawler2, error) {
	crawler, exist := t.crawlers[crawlerID]
	if !exist {
		return nil, terrors.Wrapf("Crawler %s is not found", crawlerID)
	}
	return &crawler, nil
}

func (t *CrawlerFactory2Impl) GetCrawlers(ctx context.Context, crawlerIDs ...crawler.CrawlerID) []*crawler.Crawler2 {
	crawlers := []*crawler.Crawler2{}
	for _, crawlerID := range crawlerIDs {
		crawler, exists := t.crawlers[crawlerID]
		if !exists {
			clog.L.Errorf(ctx, "crawlerID(%s) is not found", crawlerID)
			continue
		}
		crawlers = append(crawlers, &crawler)
	}
	return crawlers
}

func newCrawlerFactory2Impl(crawlers []crawler.Crawler2) *CrawlerFactory2Impl {
	factory := CrawlerFactory2Impl{
		crawlers: map[crawler.CrawlerID]crawler.Crawler2{},
	}
	for _, crawler := range crawlers {
		factory.crawlers[crawler.ID] = crawler
	}
	return &factory
}

func NewDefaultCrawlerFactory2Impl(
	repository repository.Repository,
	queue queue.Queue,
	fetcherHTTP fetcher.FetcherHTTP,
) *CrawlerFactory2Impl {
	return newCrawlerFactory2Impl(
		[]crawler.Crawler2{},
	)
}
