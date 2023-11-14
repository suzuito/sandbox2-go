package goconnpass

import (
	"net/url"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

var Def = crawler.CrawlerDefinition{
	ID: "goconnpass",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http_connpass",
		Argument: crawler.ArgumentDefinition{
			"Query": url.Values{
				"count": []string{"100"},
				"keyword_or": []string{
					"go言語",
					"golang",
					"gopher",
				},
			},
			"Days": 60,
		},
	},
	ParserDefinition: crawler.ParserDefinition{
		ID: "goconnpass",
	},
	PublisherDefinition: crawler.PublisherDefinition{
		ID: "timeseriesdatarepository",
		Argument: crawler.ArgumentDefinition{
			"ID": "goconnpass",
		},
	},
}
