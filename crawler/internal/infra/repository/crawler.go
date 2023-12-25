package repository

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/repository/internal/repositoryimpl"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewCrawlerRepository(
	defs []crawler.CrawlerDefinition,
	settings []crawler.CrawlerStarterSetting,
) repository.CrawlerRepository {
	repo := repositoryimpl.CrawlerRepository{
		Crawlers:        map[crawler.CrawlerID]*crawler.CrawlerDefinition{},
		CrawlerSettings: []*crawler.CrawlerStarterSetting{},
	}
	for i := range defs {
		repo.Crawlers[defs[i].ID] = &defs[i]
	}
	for i := range settings {
		repo.CrawlerSettings = append(repo.CrawlerSettings, &settings[i])
	}
	return &repo
}
