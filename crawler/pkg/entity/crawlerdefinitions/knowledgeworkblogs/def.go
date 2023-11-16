package knowledgeworkblogs

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

var Def = crawler.CrawlerDefinition{
	ID: "knowledgeworkblogs",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http_static",
		Argument: crawler.ArgumentDefinition{
			"URL":                "https://note.com/knowledgework/rss",
			"Method":             http.MethodGet,
			"StatusCodesSuccess": []int{http.StatusOK},
		},
	},
	ParserDefinition: crawler.ParserDefinition{
		ID: "rss",
	},
	PublisherDefinition: crawler.PublisherDefinition{
		ID: "enqueuecrawl",
		Argument: crawler.ArgumentDefinition{
			"CrawlerID": crawler.CrawlerID("knowledgeworkblog"),
		},
	},
}
