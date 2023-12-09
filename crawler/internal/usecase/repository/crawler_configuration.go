package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type CrawlerConfigurationRepository interface {
	GetDispatchCrawlSetting(ctx context.Context) (*crawler.DispatchCrawlSetting, error)
}
