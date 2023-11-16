package golangweekly

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

var Def = crawler.CrawlerDefinition{
	ID: "golangweekly",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http_static",
		Argument: crawler.ArgumentDefinition{
			"URL":                "https://cprss.s3.amazonaws.com/golangweekly.com.xml",
			"Method":             http.MethodGet,
			"StatusCodesSuccess": []int{http.StatusOK},
		},
	},
	ParserDefinition: crawler.ParserDefinition{
		ID: "rss",
	},
	PublisherDefinition: crawler.PublisherDefinition{
		ID: "timeseriesdatarepository",
		Argument: crawler.ArgumentDefinition{
			"TimeSeriesDataBaseID": timeseriesdata.TimeSeriesDataBaseID("golangweekly"),
		},
	},
}
