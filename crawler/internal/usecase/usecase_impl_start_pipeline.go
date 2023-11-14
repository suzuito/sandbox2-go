package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func (t *UsecaseImpl) StartPipelinePeriodically(
	ctx context.Context,
) error {
	t.L.Infof(ctx, "StartPipelinePeriodically")
	crawlers, err := t.CrawlerRepository.GetCrawlerDefinitions(
		ctx,
		// goblog.CrawlerID,
		// goconnpass.CrawlerID,
		// golangweekly.CrawlerID,
		// knowledgeworkblogs.CrawlerID,
	)
	if err != nil {
		return terrors.Wrap(err)
	}
	for _, crw := range crawlers {
		t.L.Infof(ctx, "Start %s (%s)", crw.ID)
		if err := t.TriggerCrawlerQueue.PublishCrawlEvent(ctx, crw.ID, crawler.CrawlerInputData{}); err != nil {
			t.L.Errorf(ctx, "PublishCrawlEvent of '%s' is failed : %+v", crw.ID, err)
			continue
		}
	}
	return nil
}
