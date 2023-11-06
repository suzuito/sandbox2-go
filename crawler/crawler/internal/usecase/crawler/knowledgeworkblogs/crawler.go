package knowledgeworkblogs

import (
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/noterss"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/queue"
)

const CrawlerID crawler.CrawlerID = "knowledgeworkblogs"

func NewCrawler(
	queue queue.Queue,
	fetcher fetcher.FetcherHTTP,
) crawler.Crawler {
	return noterss.NewCrawler(
		queue,
		fetcher,
		CrawlerID,
		"knowledgeworkblog",
		"https://note.com/knowledgework/rss",
	)
}
