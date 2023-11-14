package crawlerdefinitions

import (
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawlerdefinitions/goblog"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawlerdefinitions/goconnpass"
)

var AvailableCrawlers = []crawler.CrawlerDefinition{
	goblog.Def,
	goconnpass.Def,
}
