package repository

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/repository/internal/repositoryimpl"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewCrawlerConfigurationRepository() repository.CrawlerConfigurationRepository {
	return &repositoryimpl.CrawlerConfigurationRepository{
		Setting: &crawler.DispatchCrawlSetting{
			CrawlFunctionIDMapping: map[crawler.CrawlerID]crawler.CrawlFunctionID{
				"knowledgeworkblog": "Note",
			},
			DefaultCrawlFunctionID: "Default",
		},
	}
}
