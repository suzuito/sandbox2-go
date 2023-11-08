package knowledgeworkblogs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/queue"
	"go.uber.org/mock/gomock"
)

func TestNewCrawler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := NewCrawler(
		queue.NewMockQueue(ctrl),
		fetcher.NewMockFetcherHTTP(ctrl),
	)
	assert.Equal(t, crawler.CrawlerID("knowledgeworkblogs"), c.ID())
	assert.Equal(t, "knowledgeworkblogs", c.Name())
}
