package infra

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factorynewfuncs"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	usecase_factory "github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
)

func NewCrawlerFactory(
	setting *factorysetting.CrawlerFactorySetting,
) usecase_factory.CrawlerFactory {
	returned := factory.CrawlerFactory{
		FetcherFactory: factory.FetcherFactory{
			CrawlerFactorySetting: setting,
			NewFuncs:              factorynewfuncs.NewFuncsFetcher,
		},
		ParserFactory: factory.ParserFactory{
			CrawlerFactorySetting: setting,
			NewFuncs:              factorynewfuncs.NewFuncsParser,
		},
		PublisherFactory: factory.PublisherFactory{
			CrawlerFactorySetting: setting,
			NewFuncs:              factorynewfuncs.NewFuncsPublisher,
		},
	}
	return &returned
}
