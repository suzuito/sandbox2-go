package goblog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"go.uber.org/mock/gomock"
)

func TestCrawler(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := repository.NewMockRepository(ctrl)
	crwl := NewCrawler(repository)
	assert.Equal(t, crawler.CrawlerID("goblog"), crwl.ID())
	assert.Equal(t, "goblog", crwl.Name())
	parser, err := crwl.NewParser(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, parser)
	publisher, err := crwl.NewPublisher(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, publisher)
	fetcher, err := crwl.NewFetcher(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, fetcher)
}
