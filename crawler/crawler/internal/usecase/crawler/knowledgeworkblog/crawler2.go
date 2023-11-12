package knowledgeworkblog

import (
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/notecontent"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/publisher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

func NewCrawler2(
	fetcher crawler.Fetcher,
	repository repository.Repository,
) *crawler.Crawler2 {
	return &crawler.Crawler2{
		ID:      CrawlerID,
		Fetcher: fetcher,
		Parser: notecontent.NewParser(
			note.HasGolangTag,
		),
		Publisher: publisher.NewRepositoryToStoreTimeseriesData(
			repository,
			CrawlerID,
		),
	}
}
