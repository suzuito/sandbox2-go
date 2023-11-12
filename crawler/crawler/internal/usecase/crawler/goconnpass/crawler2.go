package goconnpass

import (
	"net/http"
	"net/url"
	"time"

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
				u, _ := url.Parse("https://connpass.com/api/v1/event/")
				q := u.Query()
				q.Add("keyword_or", "go言語")
				q.Add("keyword_or", "golang")
				q.Add("keyword_or", "gopher")
				d := time.Now()
				for i := 0; i < 30; i++ {
					q.Add("ymd", d.Add(time.Duration(i)*time.Hour*24).Format("20060102"))
				}
				q.Add("count", "100")
				u.RawQuery = q.Encode()
				r, _ := http.NewRequest(http.MethodGet, u.String(), nil)
				return r
			})(),
			func(res *http.Response) bool { return res.StatusCode == http.StatusOK },
		),
		Parser: NewParser(),
		Publisher: publisher.NewRepositoryToStoreTimeseriesData(
			repository,
			CrawlerID,
		),
	}
}
