package repositoryimpl

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type CrawlerConfigurationRepository struct {
	Setting *crawler.DispatchCrawlSetting
}

func (t *CrawlerConfigurationRepository) GetDispatchCrawlSetting(ctx context.Context) (*crawler.DispatchCrawlSetting, error) {
	return t.Setting, nil
}
