package infra

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/repositoryimpl/crawlerrepository"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewCrawlerRepository(defs []crawler.CrawlerDefinition) repository.CrawlerRepository {
	repo := crawlerrepository.Repository{
		Crawlers: map[crawler.CrawlerID]*crawler.CrawlerDefinition{},
	}
	for i := range defs {
		repo.Crawlers[defs[i].ID] = &defs[i]
	}
	return &repo
}
