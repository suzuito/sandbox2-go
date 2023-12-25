package factory

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/internal/factoryimpl"
	usecase_factory "github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
)

func NewCrawlerFactory(
	setting *factorysetting.CrawlerFactorySetting,
) usecase_factory.CrawlerFactory {
	return &factoryimpl.CrawlerFactory{
		FetcherFactory: factoryimpl.FetcherFactory{
			CrawlerFactorySetting: setting,
			NewFuncs:              factoryimpl.NewFuncsFetcher,
		},
		ParserFactory: factoryimpl.ParserFactory{
			CrawlerFactorySetting: setting,
			NewFuncs:              factoryimpl.NewFuncsParser,
		},
		PublisherFactory: factoryimpl.PublisherFactory{
			CrawlerFactorySetting: setting,
			NewFuncs:              factoryimpl.NewFuncsPublisher,
		},
	}
}
