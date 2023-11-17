package inject

import (
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawlerdefinitions/goblog"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawlerdefinitions/goconnpass"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawlerdefinitions/golangweekly"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawlerdefinitions/knowledgeworkblog"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawlerdefinitions/knowledgeworkblogs"
)

var AvailableCrawlers = []crawler.CrawlerDefinition{
	goblog.Def,
	goconnpass.Def,
	golangweekly.Def,
	knowledgeworkblogs.Def,
	knowledgeworkblog.Def,
}
