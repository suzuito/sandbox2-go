package knowledgeworkblog

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

var Def = crawler.CrawlerDefinition{
	ID: "knowledgeworkblog",
	FetcherDefinition: crawler.FetcherDefinition{
		ID: "fetcher_http",
		Argument: argument.ArgumentDefinition{
			"StatusCodesSuccess": []int{http.StatusOK},
		},
	},
	ParserDefinition: crawler.ParserDefinition{
		ID: "notecontent",
		Argument: argument.ArgumentDefinition{
			"FilterByTags": []string{
				"go",
				"golang",
				"go言語",
			},
		},
	},
	PublisherDefinition: crawler.PublisherDefinition{
		ID: "timeseriesdatarepository",
		Argument: argument.ArgumentDefinition{
			"TimeSeriesDataBaseID": timeseriesdata.TimeSeriesDataBaseID("knowledgeworkblog"),
		},
	},
}
