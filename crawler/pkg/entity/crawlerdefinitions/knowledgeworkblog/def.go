package knowledgeworkblog

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

var Def = crawler.CrawlerDefinition{
	ID: "knowledgeworkblog",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http",
		Argument: crawler.ArgumentDefinition{
			"StatusCodesSuccess": []int{http.StatusOK},
		},
	},
	ParserDefinition: crawler.ParserDefinition{
		ID: "notecontent",
		Argument: crawler.ArgumentDefinition{
			"FilterByTags": []string{
				"go",
				"golang",
				"go言語",
			},
		},
	},
	PublisherDefinition: crawler.PublisherDefinition{
		ID: "timeseriesdatarepository",
		Argument: crawler.ArgumentDefinition{
			"TimeSeriesDataBaseID": timeseriesdata.TimeSeriesDataBaseID("knowledgeworkblog"),
		},
	},
}
