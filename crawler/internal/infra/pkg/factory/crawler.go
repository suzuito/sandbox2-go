package factory

import (
	infra_internal_factory "github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	usecase_factory "github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
)

func NewCrawlerFactory(
	setting *factorysetting.CrawlerFactorySetting,
) usecase_factory.CrawlerFactory {
	return &infra_internal_factory.CrawlerFactory{
		FetcherFactory: infra_internal_factory.FetcherFactory{
			CrawlerFactorySetting: setting,
		},
		ParserFactory: infra_internal_factory.ParserFactory{
			CrawlerFactorySetting: setting,
		},
		PublisherFactory: infra_internal_factory.PublisherFactory{
			CrawlerFactorySetting: setting,
		},
	}
}
