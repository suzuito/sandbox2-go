package goconnpass

import (
	"net/url"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

var Def = crawler.CrawlerDefinition{
	ID: "goconnpass",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http_connpass",
		Argument: argument.ArgumentDefinition{
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
		Argument: argument.ArgumentDefinition{
			"TimeSeriesDataBaseID": timeseriesdata.TimeSeriesDataBaseID("goconnpass"),
		},
	},
}
