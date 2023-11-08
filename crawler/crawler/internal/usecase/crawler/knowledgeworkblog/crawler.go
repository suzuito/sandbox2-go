package knowledgeworkblog

import (
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/notecontent"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata/note"
)

const CrawlerID crawler.CrawlerID = "knowledgeworkblog"

func NewCrawler(
	repository repository.Repository,
	fetcher fetcher.FetcherHTTP,
) crawler.Crawler {
	return notecontent.NewCrawler(
		CrawlerID,
		repository,
		fetcher,
		note.HasGolangTag,
	)
}
