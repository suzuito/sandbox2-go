package repositoryimpl

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type CrawlerRepository struct {
	Crawlers        map[crawler.CrawlerID]*crawler.CrawlerDefinition
	CrawlerSettings []*crawler.CrawlerStarterSetting
}

func (t *CrawlerRepository) GetCrawlerDefinition(
	ctx context.Context,
	id crawler.CrawlerID,
) (*crawler.CrawlerDefinition, error) {
	crawlerDefinition, exists := t.Crawlers[id]
	if !exists {
		return nil, terrors.Wrapf("CrawlerDefinition[%s] is not found", id)
	}
	return crawlerDefinition, nil
}

func (t *CrawlerRepository) GetCrawlerStarterSettings(
	ctx context.Context,
	crawlerStarterSettingID crawler.CrawlerStarterSettingID,
) ([]*crawler.CrawlerStarterSetting, error) {
	returned := []*crawler.CrawlerStarterSetting{}
	for _, setting := range t.CrawlerSettings {
		if setting.ID == crawlerStarterSettingID {
			returned = append(returned, setting)
		}
	}
	return returned, nil
}
