package crawlerrepository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type Repository struct {
	Crawlers        map[crawler.CrawlerID]*crawler.CrawlerDefinition
	CrawlerSettings []*crawler.CrawlerStarterSetting
}

func (t *Repository) GetCrawlerDefinition(
	ctx context.Context,
	id crawler.CrawlerID,
) (*crawler.CrawlerDefinition, error) {
	crawlerDefinition, exists := t.Crawlers[id]
	if !exists {
		return nil, terrors.Wrapf("CrawlerDefinition[%s] is not found", id)
	}
	return crawlerDefinition, nil
}

func (t *Repository) GetCrawlerDefinitions(
	ctx context.Context,
	crawlerIDs ...crawler.CrawlerID,
) ([]*crawler.CrawlerDefinition, error) {
	defs := []*crawler.CrawlerDefinition{}
	for _, id := range crawlerIDs {
		crawlerDefinition, exists := t.Crawlers[id]
		if !exists {
			return nil, terrors.Wrapf("CrawlerDefinition[%s] is not found", id)
		}
		defs = append(defs, crawlerDefinition)
	}
	return defs, nil
}

func (t *Repository) GetCrawlerStarterSettings(
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
