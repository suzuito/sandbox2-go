package goblog

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"go.uber.org/mock/gomock"
)

func TestCrawler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := repository.NewMockRepository(ctrl)
	ftc := fetcher.NewMockFetcherHTTP(ctrl)
	crwl := NewCrawler(repository, ftc)
	assert.Equal(t, crawler.CrawlerID("goblog"), crwl.ID())
	assert.Equal(t, "goblog", crwl.Name())
}

func TestCrawlerFetch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := repository.NewMockRepository(ctrl)
	ftc := fetcher.NewMockFetcherHTTP(ctrl)
	crwl := NewCrawler(repository, ftc)

	ftc.EXPECT().DoRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	)
	crwl.Fetch(
		context.Background(),
		bytes.NewBuffer([]byte{}),
		nil,
	)
}

func TestCrawlerPublish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := repository.NewMockRepository(ctrl)
	ftc := fetcher.NewMockFetcherHTTP(ctrl)
	crwl := NewCrawler(repository, ftc)

	repository.EXPECT().SetTimeSeriesData(
		gomock.Any(),
		gomock.Any(),
	)
	crwl.Publish(
		context.Background(),
		nil,
	)
}
