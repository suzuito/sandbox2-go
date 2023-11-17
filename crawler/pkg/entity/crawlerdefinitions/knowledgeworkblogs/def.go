package knowledgeworkblogs

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/pkg/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

var Def = crawler.CrawlerDefinition{
	ID: "knowledgeworkblogs",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http_static",
		Argument: argument.ArgumentDefinition{
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
		Argument: argument.ArgumentDefinition{
			"CrawlerID": crawler.CrawlerID("knowledgeworkblog"),
		},
	},
}
