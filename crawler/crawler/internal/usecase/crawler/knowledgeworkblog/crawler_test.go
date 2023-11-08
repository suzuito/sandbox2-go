package knowledgeworkblog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"go.uber.org/mock/gomock"
)

func TestNewCrawler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := NewCrawler(
		repository.NewMockRepository(ctrl),
		fetcher.NewMockFetcherHTTP(ctrl),
	)
	assert.Equal(t, crawler.CrawlerID("knowledgeworkblog"), c.ID())
	assert.Equal(t, "knowledgeworkblog", c.Name())
}
