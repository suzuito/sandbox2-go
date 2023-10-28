package crawlerfactory

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/goblog"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/goconnpass"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/golangweekly"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
)

type CrawlerFactoryImpl struct {
	crawlers map[crawler.CrawlerID]crawler.Crawler
}

func (t *CrawlerFactoryImpl) GetCrawler(ctx context.Context, crawlerID crawler.CrawlerID) (crawler.Crawler, error) {
	crawler, exist := t.crawlers[crawlerID]
	if !exist {
		return nil, terrors.Wrapf("Crawler %s is not found", crawlerID)
	}
	return crawler, nil
}

func (t *CrawlerFactoryImpl) GetCrawlers(ctx context.Context, crawlerIDs ...crawler.CrawlerID) []crawler.Crawler {
	crawlers := []crawler.Crawler{}
	for _, crawlerID := range crawlerIDs {
		crawler, exists := t.crawlers[crawlerID]
		if !exists {
			clog.L.Errorf(ctx, "crawlerID(%s) is not found", crawlerID)
			continue
		}
		crawlers = append(crawlers, crawler)
	}
	return crawlers
}

func newCrawlerFactoryImpl(crawlers []crawler.Crawler) *CrawlerFactoryImpl {
	factory := CrawlerFactoryImpl{
		crawlers: map[crawler.CrawlerID]crawler.Crawler{},
	}
	for _, crawler := range crawlers {
		factory.crawlers[crawler.ID()] = crawler
	}
	return &factory
}

func NewDefaultCrawlerFactoryImpl(
	repository repository.Repository,
) *CrawlerFactoryImpl {
	return newCrawlerFactoryImpl(
		[]crawler.Crawler{
			goblog.NewCrawler(repository),
			goconnpass.NewCrawler(repository),
			golangweekly.NewCrawler(repository),
		},
	)
}
