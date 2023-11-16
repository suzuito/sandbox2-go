package goblog

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

var Def = crawler.CrawlerDefinition{
	ID: "goblog",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http_static",
		Argument: crawler.ArgumentDefinition{
			"URL":                "https://go.dev",
			"Method":             http.MethodGet,
			"StatusCodesSuccess": []int{http.StatusOK},
		},
	},
	ParserDefinition: crawler.ParserDefinition{
		ID: "goblog",
	},
	PublisherDefinition: crawler.PublisherDefinition{
		ID: "timeseriesdatarepository",
		Argument: crawler.ArgumentDefinition{
			"TimeSeriesDataBaseID": timeseriesdata.TimeSeriesDataBaseID("goblog"),
		},
	},
}
