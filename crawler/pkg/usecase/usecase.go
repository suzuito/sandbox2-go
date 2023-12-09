package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type Usecase interface {
	DispatchCrawlOnGCF(
		ctx context.Context,
		rawBytes []byte,
	) error
	DispatchCrawl(
		ctx context.Context,
		crawlerID crawler.CrawlerID,
		crawlerInputData crawler.CrawlerInputData,
	) error
	CrawlOnGCF(
		ctx context.Context,
		rawBytes []byte,
	) error
	Crawl(
		ctx context.Context,
		crawlerID crawler.CrawlerID,
		crawlerInputData crawler.CrawlerInputData,
	) error
	StartPipelinePeriodically(
		ctx context.Context,
		crawlerStarterSettingID crawler.CrawlerStarterSettingID,
	) error
	NotifyOnGCF(
		ctx context.Context,
		fullPath string,
	) error
}
