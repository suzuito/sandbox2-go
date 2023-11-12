package goblog

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/infra/crawlerimpl"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/publisher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
)

func NewCrawler2(
	repository repository.Repository,
) *crawler.Crawler2 {
	return &crawler.Crawler2{
		ID: CrawlerID,
		Fetcher: crawlerimpl.NewFetcherHTTPStatic(
			http.DefaultClient,
			(func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "https://go.dev", nil)
				return r
			})(),
			func(res *http.Response) bool { return res.StatusCode == http.StatusOK },
		),
		Parser: &Parser{},
		Publisher: publisher.NewRepositoryToStoreTimeseriesData(
			repository,
			CrawlerID,
		),
	}
}
