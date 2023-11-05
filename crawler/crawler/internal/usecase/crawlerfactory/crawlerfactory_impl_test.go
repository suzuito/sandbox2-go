package crawlerfactory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawler/goblog"
)

func TestCrawlerFactoryImpl(t *testing.T) {
	ctx := context.Background()
	factory := newCrawlerFactoryImpl(
		[]crawler.Crawler{
			&goblog.Crawler{},
		},
	)
	t.Run("GetCrawler", func(t *testing.T) {
		_, err := factory.GetCrawler(ctx, crawler.CrawlerID("hoge"))
		assert.Equal(t, "Crawler hoge is not found", err.Error())
		crwl, err := factory.GetCrawler(ctx, crawler.CrawlerID("goblog"))
		assert.Nil(t, err)
		assert.NotNil(t, crwl)
	})
	t.Run("GetCrawlers", func(t *testing.T) {
		crwls := factory.GetCrawlers(
			ctx,
			crawler.CrawlerID("hoge"),
			crawler.CrawlerID("goblog"),
		)
		assert.Equal(t, 1, len(crwls))
	})
	t.Run("NewDefaultCrawlerFactoryImpl", func(t *testing.T) {
		NewDefaultCrawlerFactoryImpl(nil, nil, nil)
	})
}
