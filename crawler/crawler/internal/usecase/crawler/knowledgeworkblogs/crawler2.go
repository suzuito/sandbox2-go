package knowledgeworkblogs

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/infra/crawlerimpl"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/knowledgeworkblog"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/noterss"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/publisher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/queue"
)

func NewCrawler2(
	httpClient *http.Client,
	queue queue.Queue,
) *crawler.Crawler2 {
	return &crawler.Crawler2{
		ID: CrawlerID,
		Fetcher: crawlerimpl.NewFetcherHTTPStatic(
			httpClient,
			(func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "https://note.com/knowledgework/rss", nil)
				return r
			})(),
			func(res *http.Response) bool { return res.StatusCode == http.StatusOK },
		),
		Parser: noterss.NewParser(),
		Publisher: publisher.NewPublisherToCrawler(
			queue,
			knowledgeworkblog.CrawlerID,
		),
	}
}
