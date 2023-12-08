package inject

import "github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"

var CrawlerStarterSettings = []crawler.CrawlerStarterSetting{
	{
		ID:        "starter001",
		CrawlerID: "goblog",
	},
	{
		ID:        "starter001",
		CrawlerID: "goconnpass",
	},
	{
		ID:        "starter001",
		CrawlerID: "golangweekly",
	},
	{
		ID:        "starter001",
		CrawlerID: "knowledgeworkblogs",
	},
}
