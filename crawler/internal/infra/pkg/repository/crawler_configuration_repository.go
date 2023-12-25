package repository

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/repositoryimpl/crwalerconfigurationrepository"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewCrawlerConfigurationRepository() repository.CrawlerConfigurationRepository {
	return &crwalerconfigurationrepository.CrawlerConfigurationRepository{
		Setting: &crawler.DispatchCrawlSetting{
			CrawlFunctionIDMapping: map[crawler.CrawlerID]crawler.CrawlFunctionID{
				"knowledgeworkblog": "Note",
			},
			DefaultCrawlFunctionID: "Default",
		},
	}
}
