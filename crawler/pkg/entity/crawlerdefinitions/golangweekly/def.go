package golangweekly

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

var Def = crawler.CrawlerDefinition{
	ID: "golangweekly",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http_static",
		Argument: argument.ArgumentDefinition{
			"URL":                "https://cprss.s3.amazonaws.com/golangweekly.com.xml",
			"Method":             http.MethodGet,
			"StatusCodesSuccess": []int{http.StatusOK},
			"UseCache":           false,
		},
	},
	ParserDefinition: crawler.ParserDefinition{
		ID: "rss",
	},
	PublisherDefinition: crawler.PublisherDefinition{
		ID: "timeseriesdatarepository",
		Argument: argument.ArgumentDefinition{
			"TimeSeriesDataBaseID": timeseriesdata.TimeSeriesDataBaseID("golangweekly"),
		},
	},
}
